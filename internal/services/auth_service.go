package services

import (
	"fmt"
	"time"

	"github.com/gauravst/auth-system-go/internal/config"
	"github.com/gauravst/auth-system-go/internal/models"
	"github.com/gauravst/auth-system-go/internal/repositories"
	"github.com/gauravst/auth-system-go/internal/utils/email"
	"github.com/gauravst/auth-system-go/internal/utils/hashing"
	"github.com/gauravst/auth-system-go/internal/utils/jwtToken"
	"github.com/golang-jwt/jwt/v5"
)

type AuthService interface {
	SignupUser(data *models.SignupRequest, cfg config.Config) error
	LoginUser(data *models.LoginRequest, cfg config.Config) error
	RefreshToken(token string, cfg config.Config) (string, interface{}, error)
	VerifyEmail(token string, cfg config.Config) error
	ForgotPassword(data *models.User) error
	ResetPassword(data *models.User, jwt string) error
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

func (s *authService) SignupUser(data *models.SignupRequest, cfg config.Config) error {
	// check user already exist or not
	user, err := s.authRepo.CheckUserExist(data.Username, data.Email)
	if err != nil {
		return err
	}

	if user != nil {
		return fmt.Errorf("username or email already exist")
	}

	//hash password
	hashedPassword, err := hashing.GenerateHashString(data.Password)
	if err != nil {
		return err
	}
	data.Password = hashedPassword

	// create user in db
	err = s.authRepo.SignupUser(data)
	if err != nil {
		return err
	}

	// genrete token here for email
	claims := jwt.MapClaims{
		"subject": "accountVerification",
		"email":   data.Email,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err := token.SignedString(cfg.JwtPrivateKey)

	// send email verfiction here
	toList := []string{data.Email}
	verificationURL := fmt.Sprintf("http://localhost:8080/%s", tokenString)
	body := fmt.Sprintf(
		"This is an email for account verification. Here is your verification link: %s",
		verificationURL,
	)

	smtpMail := cfg.SMTPMail
	err = email.SendEmail(smtpMail.User, smtpMail.Pass, smtpMail.Host, smtpMail.From, smtpMail.Port, toList, []byte(body))
	if err != nil {
		return err
	}

	return nil
}

func (s *authService) LoginUser(data *models.LoginRequest, cfg config.Config) error {
	// check user exist
	existUser, err := s.authRepo.CheckUserExist(data.Username, data.Email)
	if err != nil {
		return err
	}

	if existUser == nil {
		return fmt.Errorf("user not found")
	}

	// check user password
	err = hashing.CompareHashString(existUser.Password, data.Password)
	if err != nil {
		return fmt.Errorf("invalid credentials")
	}

	// check user verfied or not
	if existUser.Status != "active" {
		// send email if account not verfied
		claims := jwt.MapClaims{
			"subject": "accountVerification",
			"email":   existUser.Email,
			"exp":     time.Now().Add(24 * time.Hour).Unix(),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
		tokenString, err := token.SignedString(cfg.JwtPrivateKey)

		toList := []string{data.Email}
		verificationURL := fmt.Sprintf("http://localhost:8080/%s", tokenString)
		body := fmt.Sprintf(
			"This is an email for account verification. Here is your verification link: %s",
			verificationURL,
		)

		smtpMail := cfg.SMTPMail
		err = email.SendEmail(smtpMail.User, smtpMail.Pass, smtpMail.Host, smtpMail.From, smtpMail.Port, toList, []byte(body))
		if err != nil {
			return err
		}
		return fmt.Errorf("account not verfied, check you email")
	}

	// accessToken
	claims := jwt.MapClaims{
		"email":    existUser.Email,
		"username": existUser.Username,
		"exp":      time.Now().Add(15 * time.Minute).Unix(),
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	accessTokenString, err := accessToken.SignedString(cfg.JwtPrivateKey)
	data.AccessToken = accessTokenString

	// RefreshToken
	claims = jwt.MapClaims{
		"email":    existUser.Email,
		"username": existUser.Username,
		"exp":      time.Now().Add(24 * 30 * time.Hour).Unix(),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	refreshTokenString, err := refreshToken.SignedString(cfg.JwtPrivateKey)

	// creater new login in db
	err = s.authRepo.LoginUser(&data, refreshTokenString)
	if err != nil {
		return err
	}

	return nil
}

func (s *authService) RefreshToken(token string, cfg config.Config) (string, interface{}, error) {
	// check token
	userData, err := jwtToken.VerifyJwtAndGetData(token, cfg.JwtPrivateKey)
	if err != nil {
		if userData == nil {
			return "", nil, err
		}
	}

	// new accessToken
	jwtData := map[string]interface{}{
		"email":    userData.Email,
		"username": userData.Username,
		"exp":      time.Now().Add(15 * time.Minute).Unix(),
	}

	newToken, err := jwtToken.CreateNewToken(jwtData, cfg.JwtPrivateKey)
	if err != nil {
		return "", nil, err
	}

	// get RefreshToken from db
	refreshToken, err := s.authRepo.GetRefreshToken(userData.Email)
	if err != nil {
		return "", nil, err
	}

	// check if RefreshToken is expired or not
	_, err = jwtToken.VerifyJwtAndGetData(refreshToken, cfg.JwtPrivateKey)
	if err != nil {
		return "", nil, err
	}

	// update  jwt in db and brouwser
	return newToken, userData, nil
}

func (s *authService) VerifyEmail(token string, cfg config.Config) error {
	// check email verfiey or not
	data, err := jwtToken.VerifyJwtAndGetData(token, cfg.JwtPrivateKey)
	if err != nil {
		if data == nil {
			return err
		}
	}

	err = s.authRepo.VerifyEmail(data.email)
	if err != nil {
		return err
	}
	return nil
}

func (s *authService) ForgotPassword(data *models.User) error {
	// genrete jwt for email link token
	// send email to ResetPassword
	return nil
}

func (s *authService) ResetPassword(data *models.User, token string, cfg config.Config) (string, error) {
	// verfy jwt/ token vaild or not to ResetPassword
	jwtData, err := jwtToken.VerifyJwtAndGetData(token, cfg.JwtPrivateKey)
	if err != nil {
		return "", err
	}

	if jwtData["subject"] != "ResetPassword" {
		return "", fmt.Errorf("invalid token")
	}

	// hash new hash new password
	hashedPassword, err := hashing.GenerateHashString(data.NewPassword)
	if err != nil {
		return "", err
	}

	// save hashedPassword in db
	err = s.authRepo.ChangePassword(jwtData.Email, hashedPassword)
	if err != nil {
		return "", err
	}

	// gave jwt to login(login user) if active
	newAccessTokenData = map[string]interface{}{
		"email":    jwtData.Email,
		"username": jwtData.Username,
		"exp":      time.Now().Add(15 * time.Minute).Unix(),
	}

	accessToken, err := jwtToken.CreateNewToken(newAccessTokenData, cfg.JwtPrivateKey)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (s *authService) ChangePassword(data *models.User) error {
	// check user password right or not
	password, err := s.authRepo.CheckPassword(data)
	if err != nil {
		return err
	}

	// Compare Hashed Password
	err = hashing.CompareHashString(password, data.NewPassword)
	if err != nil {
		return err
	}

	// hash new passowrd
	hashedPassword, err := hashing.GenerateHashString(data.NewPassword)
	if err != nil {
		return err
	}

	// save hashedPassword in db
	err := s.authRepo.ChangePassword(data.Email, hashedPassword)
	if err != nil {
		return err
	}

	return nil
}
