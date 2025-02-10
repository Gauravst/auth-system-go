package middleware

import (
	"net/http"

	"github.com/gauravst/auth-system-go/internal/config"
	"github.com/gauravst/auth-system-go/internal/services"
)

func Auth(cfg *config.Config, authService *services.AuthService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Extract the token from the request headers
			cookie, err := r.Cookie("accessToken")
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			token := cookie.Value

			// refresh token
			newToken, err := authService.RefreshToken(token, cfg)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}

			// If the token is valid, call the next handler
			next.ServeHTTP(w, r)
		})
	}
}
