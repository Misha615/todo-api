package services

import (
	"todo-api/internal/domain/models"
	"todo-api/internal/domain/repositories"
)

type TaskService struct {
	repo repositories.TaskRepository
}

func NewTaskService(repo repositories.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task *models.Task) error {
	if task.Status == "" {
		task.Status = "todo"
	}
	return s.repo.Create(task)
}

func (s *TaskService) GetTasks() ([]*models.Task, error) {
	return s.repo.GetAll()
}

func (s *TaskService) GetTask(id int) ([]*models.Task, error) {
	return s.repo.GetByUserID(id)
}

func (s *TaskService) UpdateTask(task *models.Task) error {
	return s.repo.Update(task)
}

func (s *TaskService) DeleteTask(task *models.Task) error {
	return s.repo.Delete(task)
}
