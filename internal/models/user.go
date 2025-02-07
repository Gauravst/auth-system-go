package models

import (
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginSession struct {
	Id        int
	UserId    int
	Email     string
	Token     string
	IpAddress string
	Useragent string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
