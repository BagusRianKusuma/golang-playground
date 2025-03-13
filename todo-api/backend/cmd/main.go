package main

import (
	"log"
	"todo-api/internal/config"
	"todo-api/internal/handlers"
	"todo-api/internal/repositories"
	"todo-api/internal/usecases"

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

	// Define routes
	router.GET("/tasks", taskHandler.GetTasks)
	router.POST("/tasks", taskHandler.CreateTask)
	router.PUT("/tasks/:id", taskHandler.UpdateTask)
	router.POST("/tasks/:id", taskHandler.DeleteTask)

	// Start server
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
