package main

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func generateAccessToken(userID string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS512)
	claims := token.Claims.(jwt.MapClaims)
	claims["userID"] = userID
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	secretKeyString := os.Getenv("SECRET_KEY")
	tokenString, err := token.SignedString([]byte(secretKeyString))
	if err != nil {
		return "", err

	}
	return tokenString, nil
}
