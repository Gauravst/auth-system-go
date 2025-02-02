package handlers

import (
	"net/http"

	"github.com/gauravst/auth-system-go/internal/config"
	"github.com/gauravst/auth-system-go/internal/services"
	"github.com/gauravst/auth-system-go/internal/utils/response"
)

func SignupUser(authService services.AuthService, smtpMail config.SMTPMail) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.WriteJson(w, http.StatusOK, "")
	}
}

func LoginUser(authService services.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.WriteJson(w, http.StatusOK, "")
	}
}

func RefreshToken(authService services.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.WriteJson(w, http.StatusOK, "")
	}
}

func VerifyEmail(authService services.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.WriteJson(w, http.StatusOK, "")
	}
}

func ForgotPassword(authService services.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.WriteJson(w, http.StatusOK, "")
	}
}

func ResetPassword(authService services.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.WriteJson(w, http.StatusOK, "")
	}
}

func ChangePassword(authService services.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.WriteJson(w, http.StatusOK, "")
	}
}

func AuthStatus(authService services.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.WriteJson(w, http.StatusOK, "")
	}
}
