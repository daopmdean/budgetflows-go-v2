package biz

import (
	"fmt"
	"time"

	"github.com/daopmdean/budgetflows-go-v2/conf"
	"github.com/daopmdean/budgetflows-go-v2/entity"
	"github.com/golang-jwt/jwt"
)

type AppClaims struct {
	jwt.StandardClaims
	UserId   int64  `json:"userId"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func GenerateToken(user *entity.AppUser, duration time.Duration) (string, error) {
	now := time.Now()
	var claims = AppClaims{
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  now.Unix(),
			ExpiresAt: now.Add(duration).Unix(),
			Issuer:    "budgetflows.com",
		},
		UserId:   user.UserId,
		Username: user.Username,
		Email:    user.Email,
		Phone:    user.Phone,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(conf.AppConfig.SignedKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ExtractToken(token string) (*AppClaims, error) {
	var claims AppClaims

	_, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("invalid algorithm")
		}

		return []byte(conf.AppConfig.SignedKey), nil
	})
	if err != nil {
		return nil, err
	}

	err = claims.Valid()
	if err != nil {
		return nil, err
	}

	return &claims, nil
}
