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
	// fmt.Println("refresh token conformance verified")
	return nil
}

func verifyRefreshToken(refreshToken string, userID string) error {
	decodedRefreshToken, err := decodeBase64Data(refreshToken)
	if err != nil {
		return err
	}
	// fmt.Println(decodedRefreshToken)
	// isVerified := verifyData("$2a$10$rVhAhFrwPNiL5XGAcl76Au6xgCXqs3y4f/207XTdyqyYBaUn6E5QG", decodedRefreshToken)
	// fmt.Println(isVerified)
	// refreshTokenHash, err := getBscryptHash(decodedRefreshToken)
	// if err != nil {
	// 	return err
	// }
	// refreshTokenHashString := string(refreshTokenHash)
	// fmt.Println(refreshTokenHashString)
	// fmt.Println(refreshTokenHashString)
	docFromDB, err := getDocByUserID(userID)
	if err != nil {
		return errors.New("Unknown user")
	}
	refreshTokenHashFromDB := fmt.Sprintf("%v", docFromDB[1].Value)
	err = verifyHash(refreshTokenHashFromDB, decodedRefreshToken)
	if err != nil {
		return errors.New("this refresh token is invalid")
	}
	// if docFromDB[2].Value.(*int64) == nil {
	// 	return errors.New("this refresh token has expired")
	// }
	refreshTokenValidUntil, ok := docFromDB[2].Value.(int64)
	if !ok {
		return errors.New("this refresh token has expired")
	}
	if refreshTokenValidUntil < time.Now().Unix() {
		return errors.New("this refresh token has expired")
	}
	fmt.Println("refresh token verified")
	return nil
}
