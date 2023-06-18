package model

import "github.com/golang-jwt/jwt"

type AppClaims struct {
	jwt.StandardClaims
	UserId   int64  `json:"userId"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}
