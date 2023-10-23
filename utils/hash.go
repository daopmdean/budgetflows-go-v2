package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hp, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hp), err
}

func CheckPasswordHash(hash, rawPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(rawPass))
	return err == nil
}
