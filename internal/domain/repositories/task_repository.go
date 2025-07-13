package repositories

import "todo-api/internal/domain/models"

type TaskRepository interface {
	GetAll() ([]*models.Task, error)
	GetByUserID(id int) ([]*models.Task, error)
	Create(task *models.Task) error
	Update(task *models.Task) error
	Delete(task *models.Task) error
}

type UserRepository interface {
	Create(user *models.User) error
	GetByID(id int) (*models.User, error)
	Update(user *models.User) error
	Delete(user *models.User) error
}
