package main

import (
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

func getBscryptHash(data string) (string, error) {
	hashedData, err := bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)
	return string(hashedData), err
}

func encodeToBase64(data string) string {
	encodedData := base64.StdEncoding.EncodeToString([]byte(data))
	return encodedData
}

func decodeBase64Data(data string) (string, error) {
	decodedData, err := base64.StdEncoding.DecodeString(data)
	return string(decodedData), err
}

func verifyHash(verifiedHash string, dataToVerify string) error {
	err := bcrypt.CompareHashAndPassword([]byte(verifiedHash), []byte(dataToVerify))
	return err
}

func getTokenIdentifier(token string) string {
	return token[len(token)-6:]
}
