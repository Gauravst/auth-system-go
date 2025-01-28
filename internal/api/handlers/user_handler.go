package handlers

import (
	"net/http"

	"github.com/gauravst/auth-system-go/internal/services"
	"github.com/gauravst/auth-system-go/internal/utils/response"
)

func GetAllUsers(userService services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.WriteJson(w, http.StatusOK, "")
	}
}

func GetUser(userService services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.WriteJson(w, http.StatusOK, "")
	}
}

func UpdateUser(userService services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.WriteJson(w, http.StatusOK, "")
	}
}

func DeleteUser(userService services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.WriteJson(w, http.StatusOK, "")
	}
}
