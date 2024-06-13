package utils

import (
	"golang.org/x/crypto/bcrypt"
)


func HashPassword (password string) (string, error ){
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func CheckHashedPassword (password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	}
	return err == nil
}

