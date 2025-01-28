package repositories

import (
	"database/sql"

	"github.com/gauravst/auth-system-go/internal/models"
)

type AuthRepository interface {
	SignupUser(data *models.User) error
	LoginUser(data *models.User) (*models.User, error)
	RefreshToken(data *models.User) error
	VerifyEmail(data *models.User) error
	ForgotPassword(data *models.User) error
	ResetPassword(data *models.User) error
	ChangePassword(data *models.User) error
	AuthStatus(jwt string) error
}

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) SignupUser(data *models.User) error {
	return nil
}

func (r *authRepository) LoginUser(data *models.User) (*models.User, error) {
	return nil, nil
}

func (r *authRepository) RefreshToken(data *models.User) error {
	return nil
}

func (r *authRepository) VerifyEmail(data *models.User) error {
	return nil
}

func (r *authRepository) ForgotPassword(data *models.User) error {
	return nil
}

func (r *authRepository) ResetPassword(data *models.User) error {
	return nil
}

func (r *authRepository) ChangePassword(data *models.User) error {
	return nil
}

func (r *authRepository) AuthStatus(jwt string) error {
	return nil
}
