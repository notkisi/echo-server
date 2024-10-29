package util

import (
	"math/rand"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyz"

func RandomMessage(length int) string {
	rand.Seed(time.Now().UnixNano())
	message := make([]byte, length)
	for i := range message {
		message[i] = letters[rand.Intn(len(letters))]
	}
	return string(message)
}
