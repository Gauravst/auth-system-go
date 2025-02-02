package services

import (
	"github.com/gauravst/auth-system-go/internal/models"
	"github.com/gauravst/auth-system-go/internal/repositories"
)

type UserService interface {
	GetAllUsers() ([]*models.User, error)
	GetUser(id int) (*models.User, error)
	UpdateUser(id int, data *models.User) error
	DeleteUser(id int) error
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) GetAllUsers() ([]*models.User, error) {
	data, err := s.userRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *userService) GetUser(id int) (*models.User, error) {
	data, err := s.userRepo.GetUser(id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *userService) UpdateUser(id int, data *models.User) error {
	err := s.userRepo.UpdateUser(id, data)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) DeleteUser(id int) error {
	err := s.userRepo.DeleteUser(id)
	if err != nil {
		return err
	}

	return nil
}
