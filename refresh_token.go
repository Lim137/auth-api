package main

import "time"

func generateRefreshToken(accessTokenIdentifier string) (string, int64) {
	randString := generateRandomString(32)
	refreshTokenString := randString + accessTokenIdentifier
	validUntil := time.Now().Add(time.Hour * 24 * 7).Unix()
	return refreshTokenString, validUntil
}
