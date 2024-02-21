package main

import (
	"encoding/base64"
	"encoding/json"
)

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
