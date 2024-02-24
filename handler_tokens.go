package main

import (
	"fmt"
	"log"
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
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("problem with verifying the uniqueness of the GUID: %v", err.Error()))
		return
	}
	accessTokenString, err := generateAccessToken(userID)
	if err != nil {
		log.Printf("problem with access token generation: %v", err.Error())
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("problem with access token generation: %v", err.Error()))
		return
	}
	accessTokenIdentifier := getTokenIdentifier(accessTokenString)
	refreshToken, validUntil := generateRefreshToken(accessTokenIdentifier)
	hashedRefreshToken, err := getBscryptHash(refreshToken)
	if err != nil {
		log.Printf("problem with getting the hash of the token: %v", err.Error())
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("problem with getting the hash of the token: %v", err.Error()))
		return
	}
	hashedRefreshTokenString := string(hashedRefreshToken)
	err = writeRefreshTokenToDB(hashedRefreshTokenString, validUntil, userID)
	if err != nil {
		log.Printf("problem with writing to the database: %v", err.Error())
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("problem with writing to the database: %v", err.Error()))
		return

	}
	encodedRefreshToken := encodeToBase64(refreshToken)
	respond := respondingMessage{
		AccessToken:  accessTokenString,
		RefreshToken: encodedRefreshToken,
	}

	respondWithJSON(w, http.StatusOK, respond)
}
