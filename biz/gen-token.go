package biz

import (
	"fmt"
	"time"

	"github.com/daopmdean/budgetflows-go-v2/conf"
	"github.com/daopmdean/budgetflows-go-v2/entity"
	"github.com/daopmdean/budgetflows-go-v2/model"
	"github.com/daopmdean/summer/auth"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user *entity.AppUser, duration time.Duration) (string, error) {
	now := time.Now()
	claims := auth.SummerClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  &jwt.NumericDate{Time: now},
			ExpiresAt: &jwt.NumericDate{Time: now.Add(duration)},
			Issuer:    "budgetflows.com",
		},
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

func ExtractToken(token string) (*model.AppClaims, error) {
	var claims model.AppClaims
	var summerClaim auth.SummerClaim

	_, err := jwt.ParseWithClaims(token, &summerClaim, func(t *jwt.Token) (interface{}, error) {
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
