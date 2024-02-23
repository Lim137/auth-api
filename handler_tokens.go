package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

type respondingMessage struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func tokensHandler(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "guid")
	err := isUniqueUserId(userID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	accessTokenString, err := generateAccessToken(userID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	accessTokenIdentifier := getTokenIdentifier(accessTokenString)
	refreshToken, validUntil := generateRefreshToken(accessTokenIdentifier)
	fmt.Println("refresh token", refreshToken)
	hashedRefreshToken, err := getBscryptHash(refreshToken)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	hashedRefreshTokenString := string(hashedRefreshToken)
	// fmt.Println("refresh token hash", hashedRefreshTokenString)
	// isValidRefreshToken := verifyData(hashedRefreshTokenString, refreshToken)
	// if isValidRefreshToken {
	// 	fmt.Println("refresh token is valid")
	// } else {
	// 	fmt.Println("refresh token is invalid")
	// }
	err = writeRefreshTokenToDB(hashedRefreshTokenString, validUntil, userID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return

	}
	// fmt.Println("refresh token created at ", createdAt)
	encodedRefreshToken := encodeToBase64(refreshToken)
	respond := respondingMessage{
		AccessToken:  accessTokenString,
		RefreshToken: encodedRefreshToken,
	}

	respondWithJSON(w, http.StatusOK, respond)
}
