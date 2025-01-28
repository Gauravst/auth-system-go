package repositories

import (
	"database/sql"

	"github.com/gauravst/auth-system-go/internal/models"
)

type UserRepository interface {
	GetAllUsers() ([]*models.User, error)
	GetUser(id int) (*models.User, error)
	UpdateUser(id int, data *models.User) error
	DeleteUser(id int) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetAllUsers() ([]*models.User, error) {
	return nil, nil
}

func (r *userRepository) GetUser(id int) (*models.User, error) {
	return nil, nil
}

func (r *userRepository) UpdateUser(id int, data *models.User) error {
	return nil
}

func (r *userRepository) DeleteUser(id int) error {
	return nil
}
