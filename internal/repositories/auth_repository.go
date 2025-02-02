package repositories

import (
	"database/sql"
	"fmt"

	"github.com/gauravst/auth-system-go/internal/models"
)

type AuthRepository interface {
	CheckUserExist(Username, email string) (*models.User, error)
	SignupUser(data *models.User) error
	LoginUser(data *models.LoginSession) error
	RefreshToken(userId int, data *models.LoginSession) error
	VerifyEmail(id int) error
	ForgotPassword(data *models.User) error
	ResetPassword(data *models.User) error
	ChangePassword(id int, password string) error
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

func (r *userRepository) CheckUserExist(username, email string) (*models.User, error) {
	query := `
		SELECT id, username, email FROM users
		WHERE (username = $1 OR email = $2
		LIMIT 1
	`

	var user models.User
	err := r.db.QueryRow(query, username, email).Scan(&user.ID, &user.Username, &user.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *authRepository) SignupUser(data *models.User) error {
	query := `INSERT INTO users (username, email, password, status) VALUES  ($1, $2, $3, $4) RETURNING *`
	row := r.db.QueryRow(query, data.Username, data.Email, data.Password, data.Status)

	err := row.Scan(&data.ID, &data.Username, &data.Email, &data.Password, &data.Status)
	if err != nil {
		return err
	}

	return nil
}

func (r *authRepository) LoginUser(data *models.LoginSession) error {
	query := `INSERT INTO login_sessions (userId, token, ipAddress, userAgent) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(query, data.UserId, data.Token, data.IpAddress, data.Useragent)
	if err != nil {
		return err
	}

	return nil
}

func (r *authRepository) RefreshToken(userId int, data *models.LoginSession) error {
	query := `UPDATE login_sessions SET token = $1, ipAddress = $2, userAgent = $3  status = $4 WHERE userId = $5`
	_, err := r.db.Exec(query, data.Token, data.IpAddress, data.Useragent, data.Status, userId)
	if err != nil {
		return err
	}

	return nil
}

func (r *authRepository) VerifyEmail(id int) error {
	query := `UPDATE users SET status = 'active' WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *authRepository) ForgotPassword(data *models.User) error {
	// wait
	return nil
}

func (r *authRepository) ResetPassword(data *models.User) error {
	// wait
	return nil
}

func (r *authRepository) ChangePassword(id int, password string) error {
	query := `UPDATE users SET password = $1 where id = $2`
	_, err := r.db.Exec(query, password, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *authRepository) AuthStatus(jwt string) error {
	var status string

	query := `SELECT status FROM login_sessions WHERE token = $1`
	err := r.db.QueryRow(query, jwt).Scan(&status)
	if err != nil {
		return err
	}

	if status != "active" {
		return fmt.Errorf("auth status not active %s", status)
	}

	return nil
}
