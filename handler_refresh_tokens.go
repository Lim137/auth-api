package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func refreshTokensHandler(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	defer r.Body.Close()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	err = verifyTokenConformance(params.AccessToken, params.RefreshToken)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	partsOfAccessToken := strings.Split(params.AccessToken, ".")
	accessTokenClaims, err := decodeAccessTokenClaims(partsOfAccessToken[1])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	userID := fmt.Sprintf("%v", accessTokenClaims["userID"])
	err = verifyRefreshToken(params.RefreshToken, userID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	// Создать новые токены и отдать их + заменить refresh токен в БД
	newAccessTokenString, err := generateAccessToken(userID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	accessTokenIdentifier := getTokenIdentifier(newAccessTokenString)
	newRefreshToken, newTokenValidUntil := generateRefreshToken(accessTokenIdentifier)
	// fmt.Println("refresh token", newRefreshToken)
	hashedNewRefreshToken, err := getBscryptHash(newRefreshToken)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	hashedNewRefreshTokenString := string(hashedNewRefreshToken)
	err = updateRefreshTokenInDB(userID, hashedNewRefreshTokenString, newTokenValidUntil)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	encodedNewRefreshToken := encodeToBase64(newRefreshToken)
	respond := respondingMessage{
		AccessToken:  newAccessTokenString,
		RefreshToken: encodedNewRefreshToken,
	}

	respondWithJSON(w, http.StatusOK, respond)
	// fmt.Println(params.AccessToken)
	// fmt.Println(params.RefreshToken)
}
