package services

import (
	"github.com/gauravst/auth-system-go/internal/models"
	"github.com/gauravst/auth-system-go/internal/repositories"
)

type AuthService interface {
	SignupUser(data *models.User) error
	LoginUser(data *models.User) (*models.User, error)
	RefreshToken(data *models.User) error
	VerifyEmail(data *models.User) error
	ForgotPassword(data *models.User) error
	ResetPassword(data *models.User) error
	ChangePassword(data *models.User) error
	AuthStatus(jwt string) error
}

type authService struct {
	authRepo repositories.AuthRepository
}

func NewAuthService(authRepo repositories.AuthRepository) AuthService {
	return &authService{
		authRepo: authRepo,
	}
}

func (s *authService) SignupUser(data *models.User) error {
	return nil
}

func (s *authService) LoginUser(data *models.User) (*models.User, error) {
	return nil, nil
}

func (s *authService) RefreshToken(data *models.User) error {
	return nil
}

func (s *authService) VerifyEmail(data *models.User) error {
	return nil
}

func (s *authService) ForgotPassword(data *models.User) error {
	return nil
}

func (s *authService) ResetPassword(data *models.User) error {
	return nil
}

func (s *authService) ChangePassword(data *models.User) error {
	return nil
}

func (s *authService) AuthStatus(data *models.User) error {
	return nil
}
