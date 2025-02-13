package middleware

import (
	"context"
	"net/http"

	"github.com/gauravst/auth-system-go/internal/config"
	"github.com/gauravst/auth-system-go/internal/services"
	"github.com/gauravst/auth-system-go/internal/utils/jwtToken"
	"github.com/gauravst/auth-system-go/internal/utils/response"
)

type contextKey string

const userDataKey contextKey = "userData"

func Auth(cfg *config.Config, authService *services.AuthService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Extract the token from the request headers
			cookie, err := r.Cookie("accessToken")
			if err != nil {
				response.WriteJson(w, http.StatusUnauthorized, response.GeneralError(err))
				return
			}
			token := cookie.Value

			// refresh token
			newToken, data, err := authService.RefreshToken(token, cfg)
			if err != nil {
				response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
				return
			}

			jwtToken.SetAccessToken(w, newToken, false)

			ctx := context.WithValue(r.Context(), userDataKey, data)
			r.WithContext(ctx)

			// If the token is valid, call the next handler
			next.ServeHTTP(w, r)
		})
	}
}

func AuthOnlyAdmin(cfg *config.Config, authService *services.AuthService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			cookie, err := r.Cookie("accessToken")
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			token := cookie.Value

			newToken, data, err := authService.RefreshToken(token, cfg)
			if err != nil {
				response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
				return
			}

			jwtToken.SetAccessToken(w, newToken, false)

			if data.Role != "ADMIN" {
				response.WriteJson(w, http.StatusUnauthorized, response.GeneralError(err))
				return
			}

			ctx := context.WithValue(r.Context(), userDataKey, data)
			r.WithContext(ctx)

			// If the token is valid, call the next handler
			next.ServeHTTP(w, r)
		})
	}
}
