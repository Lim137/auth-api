package main

import (
	"encoding/base64"
	"encoding/json"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func generateAccessToken(userID string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS512)
	claims := token.Claims.(jwt.MapClaims)
	claims["userID"] = userID
	claims["exp"] = time.Now().Add(time.Minute * 2).Unix()
	secretKeyString := os.Getenv("SECRET_KEY")
	tokenString, err := token.SignedString([]byte(secretKeyString))
	if err != nil {
		return "", err

	}
	return tokenString, nil
}

func decodeAccessTokenClaims(payloadString string) (map[string]interface{}, error) {
	payload, err := base64.RawURLEncoding.DecodeString(payloadString)
	if err != nil {
		return nil, err
	}
	var tokenClaims map[string]interface{}
	err = json.Unmarshal(payload, &tokenClaims)
	if err != nil {
		return nil, err
	}
	return tokenClaims, nil
}
