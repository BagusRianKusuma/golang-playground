package main

import (
	"log"
	"todo-api/internal/config"
	"todo-api/internal/handlers"
	"todo-api/internal/repositories"
	"todo-api/internal/usecases"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to database
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize repository, use case, and handler
	taskRepo := repositories.NewTaskRepository(db)
	taskUseCase := usecases.NewTaskUseCase(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskUseCase)

	// Setup Gin router
	router := gin.Default()

	//middleware CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5500", "http://127.0.0.1:5500"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	}))

	// Define routes
	router.GET("/tasks", taskHandler.GetTasks)
	router.POST("/tasks", taskHandler.CreateTask)
	router.PUT("/tasks/:id", taskHandler.UpdateTask)
	router.DELETE("/tasks/:id", taskHandler.DeleteTask)

	// Start server
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
