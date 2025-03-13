package usecases

import (
	"todo-api/internal/entities"
	"todo-api/internal/repositories"
)

type TaskUseCase struct {
	taskRepo *repositories.TaskRepository
}

func NewTaskUseCase(taskRepo *repositories.TaskRepository) *TaskUseCase {
	return &TaskUseCase{taskRepo: taskRepo}
}

func (uc *TaskUseCase) CreateTask(title, description string) (*entities.Task, error) {
	task := &entities.Task{
		Title:       title,
		Description: description,
		Completed:   false,
	}
	if err := uc.taskRepo.CreateTask(task); err != nil {
		return nil, err
	}
	return task, nil
}

func (uc *TaskUseCase) GetTask() ([]entities.Task, error) {
	return uc.taskRepo.GetTask()
}

func (uc *TaskUseCase) UpdateTask(id int, title, description string, completed bool) error {
	task := &entities.Task{
		Title:       title,
		Description: description,
		Completed:   completed,
	}
	return uc.taskRepo.UpdateTask(id, task)
}

func (uc *TaskUseCase) DeleteTask(id int) error {
	return uc.taskRepo.DeleteTask(id)
}
