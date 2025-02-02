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
	query := `SELECT id, username , email, password, status FROM users`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []*models.User
	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Status)
		if err != nil {
			return nil, err
		}

		data = append(data, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

func (r *userRepository) GetUser(id int) (*models.User, error) {
	data := &models.User{}
	query := `SELECT id, username, email, password, status FROM users WHERE id = $1`
	row := r.db.QueryRow(query, id)

	err := row.Scan(
		&data.ID,
		&data.Username,
		&data.Email,
		&data.Password,
		&data.Status,
	)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *userRepository) UpdateUser(id int, data *models.User) error {
	query := `UPDATE users SET username = $1, email = $2 , password = $3 WHERE id = $4 RETURNING id, username , email, password, status`
	row := r.db.QueryRow(query, data.Username, data.Email, data.Password, id)

	err := row.Scan(&data.ID, &data.Username, &data.Email, &data.Password, &data.Status)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) DeleteUser(id int) error {
	query := `DELETE users WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
