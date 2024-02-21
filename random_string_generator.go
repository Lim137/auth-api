package main

import (
	"math/rand"
	"time"
)

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var randNumGenerator *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, length)
	for i := range bytes {
		bytes[i] = charset[randNumGenerator.Intn(len(charset))]

	}
	return string(bytes)
}
