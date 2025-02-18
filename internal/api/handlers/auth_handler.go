package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gauravst/auth-system-go/internal/config"
	"github.com/gauravst/auth-system-go/internal/models"
	"github.com/gauravst/auth-system-go/internal/services" "github.com/gauravst/auth-system-go/internal/utils/jwtToken"
	"github.com/gauravst/auth-system-go/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func SignupUser(authService services.AuthService, cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data models.SignupRequest
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			response.WriteJson(w, http.StatusNotFound, response.GeneralError(err))
			return
		}

		err = validator.New().Struct(data)
		if err != nil {
			validateErrs, ok := err.(validator.ValidationErrors)
			if ok {
				response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrs))
				return
			}

			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("Bad Request")))
			return
		}

		err = authService.SignupUser(&data, *cfg)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		response.WriteJson(w, http.StatusOK, "Account created, check you email to verify account")
		return
	}
}

func LoginUser(authService services.AuthService, cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get data for Request
		var data models.LoginRequest
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			response.WriteJson(w, http.StatusNotFound, response.GeneralError(err))
			return
		}

		// validat data
		err = validator.New().Struct(data)
		if err != nil {
			validateErrs, ok := err.(validator.ValidationErrors)
			if ok {
				response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrs))
				return
			}

			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("Bad Request")))
			return
		}

		// login user
		err = authService.LoginUser(&data)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		//set login cookies
		if data.AccessToken != " " {
			jwtToken.SetAccessToken(w, data.AccessToken, false)
		}

		// return login data
		response.WriteJson(w, http.StatusOK, data)
		return
	}
}

func RefreshToken(authService services.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// work done here becouse new token by middware done
		response.WriteJson(w, http.StatusOK, "")
		return
	}
}

func VerifyEmail(authService services.AuthService, cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.PathValue("token")
		if token == " " {
			response.WriteJson(w, http.StatusNotFound, response.GeneralError(fmt.Errorf("token parms not found")))
			return
		}

		err := authService.VerifyEmail(token, *cfg)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		response.WriteJson(w, http.StatusOK, "Account verifed")
		return
	}
}

func ForgotPassword(authService services.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get data from Request
		var data models.ForgotPasswordRequest
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			response.WriteJson(w, http.StatusNotFound, response.GeneralError(err))
			return
		}

		// vaildate data
		err = validator.New().Struct(data)
		if err != nil {
			validateErrs, ok := err.(validator.ValidationErrors)
			if ok {
				response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrs))
				return
			}

			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("Bad Request")))
			return
		}

		// ForgotPassword
		err := authService.ForgotPassword(&data)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		// return response
		response.WriteJson(w, http.StatusOK, "Reset password link send to you email")
		return
	}
}

func ResetPassword(authService services.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get parmas
		token := r.PathValue("token")
		if token == " " {
			response.WriteJson(w, http.StatusNotFound, response.GeneralError(fmt.Errorf("token parms not found")))
			return
		}

		// get data
		var data models.ResetPasswordRequest
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			response.WriteJson(w, http.StatusNotFound, response.GeneralError(err))
			return
		}

		// vaildate data
		err = validator.New().Struct(data)
		if err != nil {
			validateErrs, ok := err.(validator.ValidationErrors)
			if ok {
				response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrs))
				return
			}

			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("Bad Request")))
			return
		}

		// ResetPassword with login
		err = authService.ResetPassword(&data, token)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		// return response
		response.WriteJson(w, http.StatusOK, "Password Updated")
		return
	}
}

func ChangePassword(authService services.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get data
		var data models.ChangePasswordRequest
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			response.WriteJson(w, http.StatusNotFound, response.GeneralError(err))
			return
		}

		// vaildate data
		err = validator.New().Struct(data)
		if err != nil {
			validateErrs, ok := err.(validator.ValidationErrors)
			if ok {
				response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrs))
				return
			}

			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("Bad Request")))
			return
		}

		// ResetPassword with login
		err = authService.ChangePassword(&data)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		// return response
		response.WriteJson(w, http.StatusOK, "Password changed")
		return
	}
}

func AuthStatus(authService services.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//work done by auth middware
		var userData models.User
		userData, ok := r.Context().Value(userDataKey).(models.User)
		if !ok {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(fmt.Errorf("user data not found")))
			return
		}

		response.WriteJson(w, http.StatusOK, userData)
		return
	}
}
