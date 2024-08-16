package services

import (
	"github.com/printSANO/go-gin-example/models"
	"github.com/printSANO/go-gin-example/repositories"
)

type UserService interface {
	GetUserByID(id uint) (*models.User, error)
	CreateUser(user *models.User) error
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) CreateUser(user *models.User) error {
	return s.repo.Create(user)
}
