package utils

import (
	"math/rand"
	"time"
)

func GenerateRandomString(length int) string {
	// Define the alphabet
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Initialize an empty string to store the result
	result := make([]byte, length)

	// Generate random characters
	for i := 0; i < length; i++ {
		result[i] = alphabet[rand.Intn(len(alphabet))]
	}

	// Convert result to string and return
	return string(result)
}
