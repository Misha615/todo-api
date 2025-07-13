package services

import (
	"todo-api/internal/domain/models"
	"todo-api/internal/domain/repositories"
)

type UserService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.repo.Create(user)
}

func (s *UserService) GetUser(id int) (*models.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) UpdateUser(user *models.User) error {
	return s.repo.Update(user)
}

func (s *UserService) DeleteUser(user *models.User) error {
	return s.repo.Delete(user)
}
