package main

import (
	"errors"
	"fmt"
	"time"
)

func verifyTokenConformance(accessToken string, refreshToken string) error {
	decodedRefreshToken, err := decodeBase64Data(refreshToken)
	if err != nil {
		return err
	}
	refreshTokenIdentifier := getTokenIdentifier(decodedRefreshToken)
	accessTokenIdentifier := getTokenIdentifier(accessToken)
	if refreshTokenIdentifier != accessTokenIdentifier {
		return errors.New("this refresh token is for another access token")
	}
	return nil
}

func verifyRefreshToken(refreshToken string, userID string) error {
	decodedRefreshToken, err := decodeBase64Data(refreshToken)
	if err != nil {
		return err
	}
	docFromDB, err := getDocByUserID(userID)
	if err != nil {
		return errors.New("Unknown user")
	}
	refreshTokenHashFromDB := fmt.Sprintf("%v", docFromDB[1].Value)
	err = verifyHash(refreshTokenHashFromDB, decodedRefreshToken)
	if err != nil {
		return errors.New("this refresh token is invalid")
	}
	refreshTokenValidUntil, ok := docFromDB[2].Value.(int64)
	if !ok {
		return errors.New("this refresh token has expired")
	}
	if refreshTokenValidUntil < time.Now().Unix() {
		return errors.New("this refresh token has expired")
	}
	return nil
}
