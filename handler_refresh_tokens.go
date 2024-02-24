package main

import (
	"encoding/json"
	"fmt"
	"log"
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
		log.Printf("couldn't decode body parameters: %v", err.Error())
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("couldn't decode body parameters: %v", err.Error()))
		return
	}
	err = verifyTokenConformance(params.AccessToken, params.RefreshToken)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("couldn't confirm the tokens match: %v", err.Error()))
		return
	}

	partsOfAccessToken := strings.Split(params.AccessToken, ".")
	accessTokenClaims, err := decodeAccessTokenClaims(partsOfAccessToken[1])
	if err != nil {
		log.Printf("couldn't decode access token claims: %v", err.Error())
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("couldn't decode access token claims: %v", err.Error()))
		return
	}
	userID := fmt.Sprintf("%v", accessTokenClaims["userID"])
	err = verifyRefreshToken(params.RefreshToken, userID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("couldn't verify refresh token: %v", err.Error()))
		return
	}
	newAccessTokenString, err := generateAccessToken(userID)
	if err != nil {
		log.Printf("couldn't generate new access token: %v", err.Error())
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("couldn't generate new access token: %v", err.Error()))
		return
	}
	accessTokenIdentifier := getTokenIdentifier(newAccessTokenString)
	newRefreshToken, newTokenValidUntil := generateRefreshToken(accessTokenIdentifier)
	hashedNewRefreshToken, err := getBscryptHash(newRefreshToken)
	if err != nil {
		log.Printf("couldn't get the hash of the token: %v", err.Error())
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("couldn't get the hash of the token: %v", err.Error()))
		return
	}
	hashedNewRefreshTokenString := string(hashedNewRefreshToken)
	err = updateRefreshTokenInDB(userID, hashedNewRefreshTokenString, newTokenValidUntil)
	if err != nil {
		log.Printf("couldn't update the data in the database: %v", err.Error())
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("couldn't update the data in the database: %v", err.Error()))
		return
	}
	encodedNewRefreshToken := encodeToBase64(newRefreshToken)
	respond := respondingMessage{
		AccessToken:  newAccessTokenString,
		RefreshToken: encodedNewRefreshToken,
	}

	respondWithJSON(w, http.StatusOK, respond)
}
