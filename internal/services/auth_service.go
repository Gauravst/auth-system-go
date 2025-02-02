package services

import (
	"fmt"
	"time"

	"github.com/gauravst/auth-system-go/internal/config"
	"github.com/gauravst/auth-system-go/internal/models"
	"github.com/gauravst/auth-system-go/internal/repositories"
	"github.com/gauravst/auth-system-go/internal/utils/email"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	SignupUser(data *models.User, smtpMail config.SMTPMail) error
	LoginUser(data *models.LoginSession) (*models.User, error)
	RefreshToken(data *models.LoginSession) error
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

func (s *authService) SignupUser(data *models.User, smtpMail config.SMTPMail) error {
	// check user already exist or not
	user, err := s.authRepo.CheckUserExist(data.Username, data.Email)
	if err != nil {
		return err
	}

	if user != nil {
		return fmt.Errorf("username or email already exist")
	}

	//hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	data.Password = string(hashedPassword)

	// create user in db
	err = s.authRepo.SignupUser(data)
	if err != nil {
		return err
	}

	// genrete token here for email
	claims := jwt.MapClaims{
		"email": data.Email,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	// send email verfiction here
	toList := []string{data.Email}
	verificationURL := fmt.Sprintf("http://localhost:8080/api/auth/verify/%s", token)
	body := fmt.Sprintf(
		"This is an email for account verification. Here is your verification link: %s",
		verificationURL,
	)

	err = email.SendEmail(smtpMail.User, smtpMail.Pass, smtpMail.Host, smtpMail.From, smtpMail.Port, toList, []byte(body))
	if err != nil {
		return err
	}

	return nil
}

func (s *authService) LoginUser(data *models.LoginSession) (*models.User, error) {
	// check user exist
	// check user password
	// check user verfied or not
	// genrete user login jwt
	// send back using http (handler)
	return nil, nil
}

func (s *authService) RefreshToken(data *models.LoginSession) error {
	// check user login session
	// update  jwt in db and brouwser
	return nil
}

func (s *authService) VerifyEmail(data *models.User) error {
	// check email verfiey or not
	// send email
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

func (s *authService) AuthStatus(jwt string) error {
	return nil
}
