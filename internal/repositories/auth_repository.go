package repositories

import (
	"database/sql"

	"github.com/gauravst/auth-system-go/internal/models"
)

type AuthRepository interface {
	CheckUserExist(username, email string) (*models.User, error)
	SignupUser(data *models.SignupRequest) error
	LoginUser(data *models.LoginRequest, refreshToken string) error
	GetRefreshToken(email string) (string, error)
	VerifyEmail(email string) error
	ForgotPassword(data *models.ForgotPasswordRequest) error
	ResetPassword(data *models.ResetPasswordRequest) error
	ChangePassword(email, password string) error
}

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) CheckUserExist(username, email string) (*models.User, error) {
	query := `
		SELECT id, username, email FROM users
    WHERE (username = $1 OR email = $2) LIMIT 1
	`

	var user *models.User
	err := r.db.QueryRow(query, username, email).Scan(&user.ID, &user.Username, &user.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func (r *authRepository) SignupUser(data *models.SignupRequest) error {
	query := `INSERT INTO users (username, email, password) VALUES  ($1, $2, $3) RETURNING *`
	row := r.db.QueryRow(query, data.Username, data.Email, data.Password)

	err := row.Scan(&data.Id, &data.Username, &data.Email, &data.Password, &data.Status)
	if err != nil {
		return err
	}

	return nil
}

func (r *authRepository) LoginUser(data *models.LoginRequest, refreshToken string) error {
	query := `INSERT INTO login_sessions (userId, token ) VALUES ($1, $2 )`
	_, err := r.db.Exec(query, data.UserId, refreshToken)
	if err != nil {
		return err
	}

	return nil
}

func (r *authRepository) GetRefreshToken(email string) (string, error) {
	var token string

	query := `SELECT token FROM login_sessions WHERE email = $1`
	err := r.db.QueryRow(query, email).Scan(&token)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *authRepository) VerifyEmail(email string) error {
	query := `UPDATE users SET status = 'active' WHERE email = $1`
	_, err := r.db.Exec(query, email)
	if err != nil {
		return err
	}

	return nil
}

func (r *authRepository) ForgotPassword(data *models.ForgotPasswordRequest) error {
	// wait
	return nil
}

func (r *authRepository) ResetPassword(data *models.ResetPasswordRequest) error {
	// wait
	return nil
}

func (r *authRepository) ChangePassword(email, password string) error {
	query := `UPDATE users SET password = $1 where email = $2`
	_, err := r.db.Exec(query, password, email)
	if err != nil {
		return err
	}

	return nil
}
