package main

import "time"

func generateRefreshToken(accessTokenIdentifier string) (string, int64) {
	randString := generateRandomString(32)
	refreshTokenString := randString + accessTokenIdentifier
	createdAt := time.Now().Unix()
	return refreshTokenString, createdAt
}
