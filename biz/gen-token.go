package biz

import (
	"time"

	"github.com/daopmdean/budgetflows-go-v2/conf"
	"github.com/daopmdean/budgetflows-go-v2/entity"
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
		UserId:   user.UserId,
		Username: user.Username,
		Email:    user.Email,
		Phone:    user.Phone,
	}

	signedToken, err := auth.GenToken(claims, conf.AppConfig.SignedKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
