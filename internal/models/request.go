package models

import "time"

type SignupRequest struct {
	Id       int    `json:"id,omitempty"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Status   string `json:"status,omitempty"`
}

type LoginRequest struct {
	Id          int       `json:"id,omitempty"`
	Username    string    `json:"username,omitempty"`
	Email       string    `json:"email", validate:"required"`
	Password    string    `json:"password" validate:"required"`
	Status      string    `json:"status, omitempty"`
	AccessToken string    `json:"accessToken, omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ForgotPasswordRequest struct {
	Id    int    `json:"id, omitempty"`
	Email string `json:"email" validate : "required"`
}

type ChangePasswordRequest struct {
	Email       string `json:"email" validate : "required"`
	Password    string `json:"password, omitempty" validate : "required"`
	NewPassword string `json:"NewPassword, omitempty" validate : "required"`
}

type ResetPasswordRequest struct {
	Email    string `json:"email" validate : "required"`
	Password string `json:"password, omitempty" validate : "required"`
}
