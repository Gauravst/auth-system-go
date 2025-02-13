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
	"github.com/gauravst/auth-system-go/internal/api/middleware"
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

	// user router/ routes
	userRepo := repositories.NewUserRepository(database.DB)
	userService := services.NewUserService(userRepo)

	// auth router/ routes
	authRepo := repositories.NewAuthRepository(database.DB)
	authService := services.NewAuthService(authRepo)

	// get all users data
	router.Handle("GET /api/users",
		middleware.Auth(cfg, &authService)(
			http.HandlerFunc(handlers.GetAllUsers(userService)),
		),
	)

	// get user data currently login
	router.Handle("GET /api/user",
		middleware.Auth(cfg, &authService)(
			http.HandlerFunc(handlers.GetUser(userService)),
		),
	)

	//  only admin can do ..
	router.Handle("GET /api/user/{id}",
		middleware.Auth(cfg, &authService)(
			http.HandlerFunc(handlers.GetUserById(userService)),
		),
	)

	router.Handle("PUT /api/user/{id}",
		middleware.Auth(cfg, &authService)(
			http.HandlerFunc(handlers.UpdateUser(userService)),
		),
	)

	router.Handle("DELETE /api/user/{id}",
		middleware.Auth(cfg, &authService)(
			http.HandlerFunc(handlers.DeleteUser(userService)),
		),
	)

	router.HandleFunc("POST /api/auth/signup", handlers.SignupUser(authService, cfg))
	router.HandleFunc("POST /api/auth/login", handlers.LoginUser(authService, cfg))

	// this is temp.. to check middleware
	router.Handle("POST /api/auth/refresh",
		middleware.Auth(cfg, &authService)(
			http.HandlerFunc(handlers.RefreshToken(authService)),
		),
	)

	// resend verification mail to verify user account
	router.HandleFunc("POST /api/auth/resend-verification", handlers.VerifyEmail(authService))

	// verify email using link
	router.HandleFunc("GET /api/{token}", handlers.VerifyEmail(authService, cfg))

	// if user forgot password and user want to reset password using email
	router.HandleFunc("POST /api/auth/forgot-password", handlers.ForgotPassword(authService))

	// user want reset password and want set new password (wothout old password)
	router.HandleFunc("POST /api/auth/reset-password", handlers.ResetPassword(authService))

	// change user password with old(current) and new password
	router.HandleFunc("POST /api/auth/change-password", handlers.ChangePassword(authService))

	// to check user login or logut
	router.Handle("POST /api/auth/status",
		middleware.Auth(cfg, &authService)(
			http.HandlerFunc(handlers.AuthStatus(authService)),
		),
	)

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
		slog.Error("failed to Shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("server Shutdown successfully")
}
