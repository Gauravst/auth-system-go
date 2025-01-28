package middleware

import (
	"net/http"
)

func Auth(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the token from the request headers
		token := r.Header.Get("Authorization")

		if !isValidToken(token) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// If the token is valid, call the next handler
		next.ServeHTTP(w, r)
	})
}

func isValidToken(token string) bool {
	// In a real application, you would validate the token against a database or a JWT library
	// For this example, we'll assume the token is valid if it's not empty
	return token != ""
}
