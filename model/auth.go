package model

import "time"

type LoginRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`

	Name    string     `json:"name"`
	Address string     `json:"address"`
	Avatar  string     `json:"avatar"`
	Dob     *time.Time `json:"dob"`
}
