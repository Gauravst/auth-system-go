package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gauravst/auth-system-go/internal/api/handlers"
	"github.com/gauravst/auth-system-go/internal/config"
	"github.com/gauravst/auth-system-go/internal/database"
	"github.com/gauravst/auth-system-go/internal/repositories"
	"github.com/gauravst/auth-system-go/internal/services"
)

func main() {
	// load config
	cfg := config.ConfigMustLoad()

	// database setup
	database.InitDB(cfg.DatabaseUri)
	defer database.CloseDB()

	//setup router
	router := http.NewServeMux()

	userRepo := repositories.NewUserRepository(database.DB)
	userService := services.NewUserService(userRepo)

	router.HandleFunc("GET /api/user", handlers.GetAllUsers(userService))
	router.HandleFunc("GET /api/user/{id}", handlers.GetUser(userService))
	router.HandleFunc("PUT /api/user/{id}", handlers.UpdateUser(userService))
	router.HandleFunc("DELETE /api/user/{id}", handlers.DeleteUser(userService))

	authRepo := repositories.NewAuthRepository(database.DB)
	authService := services.NewAuthService(authRepo)

	router.HandleFunc("POST /api/auth/signup", handlers.SignupUser(authService))
	router.HandleFunc("POST /api/auth/login", handlers.LoginUser(authService))
	router.HandleFunc("POST /api/auth/refresh", handlers.RefreshToken(authService))
	router.HandleFunc("POST /api/auth/resend-verification", handlers.VerifyEmail(authService))
	router.HandleFunc("POST /api/auth/forgot-password", handlers.ForgotPassword(authService))
	router.HandleFunc("POST /api/auth/reset-password", handlers.ResetPassword(authService))
	router.HandleFunc("POST /api/auth/change-password", handlers.ChangePassword(authService))
	router.HandleFunc("GET /api/auth/status", handlers.AuthStatus(authService))

	// setup server
	server := &http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}

	slog.Info("server started", slog.String("address", cfg.Address))

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("failed to start server")
		}
	}()

	<-done

	slog.Info("shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("faild to Shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("server Shutdown successfully")
}
