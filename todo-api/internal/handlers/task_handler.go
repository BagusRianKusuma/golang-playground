package handlers

import (
	"todo-api/internal/entities"
	"todo-api/internal/usecases"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	taskUseCase *usecases.TaskUseCase
}

func NewTaskHandler(taskUseCase *usecases.TaskUseCase) *TaskHandler {
	return &TaskHandler{taskUseCase: taskUseCase}
}

func (h *TaskHandler) GetTasks(c *gin.Context) {
	tasks, err := h.taskUseCase.GetTask()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, tasks)
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var task entities.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	createdTask, err := h.taskUseCase.CreateTask(task.Title, task.Description)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, createdTask)
}
