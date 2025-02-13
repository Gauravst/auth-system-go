package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gauravst/auth-system-go/internal/models"
	"github.com/gauravst/auth-system-go/internal/services"
	"github.com/gauravst/auth-system-go/internal/utils/response"
)

type contextKey string

const userDataKey contextKey = "userData"

func GetAllUsers(userService services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		data, err := userService.GetAllUsers()
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		response.WriteJson(w, http.StatusOK, data)
		return
	}
}

func GetUser(userService services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userData, ok := r.Context().Value(userDataKey).(interface{})
		if !ok {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(fmt.Errorf("user data not found")))
			return
		}

		response.WriteJson(w, http.StatusOK, userData)
		return
	}
}

func GetUserById(userService services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if id == " " {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(fmt.Errorf("id parms not found")))
			return
		}

		userData, ok := r.Context().Value(userDataKey).(interface{})
		if !ok {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(fmt.Errorf("user data not found")))
			return
		}

		idInt, err := strconv.Atoi(id)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		if userData.role != "ADMIN" && userData.Id != idInt {
			response.WriteJson(w, http.StatusUnauthorized, response.GeneralError(fmt.Errorf("unauthorized user")))
			return
		}

		data, err := userService.GetUser(idInt)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		response.WriteJson(w, http.StatusOK, data)
		return
	}
}

func UpdateUser(userService services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if id == " " {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(fmt.Errorf("id parms not found")))
			return
		}

		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			response.WriteJson(w, http.StatusNotFound, response.GeneralError(err))
			return
		}

		userData, ok := r.Context().Value(userDataKey).(interface{})
		if !ok {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(fmt.Errorf("user data not found")))
			return
		}

		idInt, err := strconv.Atoi(id)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		if userData.role != "ADMIN" && userData.Id != idInt {
			response.WriteJson(w, http.StatusUnauthorized, response.GeneralError(fmt.Errorf("unauthorized user")))
			return
		}

		err = userService.UpdateUser(idInt, &user)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		response.WriteJson(w, http.StatusOK, user)
		return
	}
}

func DeleteUser(userService services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if id == " " {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(fmt.Errorf("id parms not found")))
			return
		}

		userData, ok := r.Context().Value(userDataKey).(interface{})
		if !ok {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(fmt.Errorf("user data not found")))
			return
		}

		idInt, err := strconv.Atoi(id)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		if userData.role != "ADMIN" && userData.Id != idInt {
			response.WriteJson(w, http.StatusUnauthorized, response.GeneralError(fmt.Errorf("unauthorized user")))
			return
		}

		err = userService.DeleteUser(idInt)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		response.WriteJson(w, http.StatusOK, "user deleted")
		return
	}
}
