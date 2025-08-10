package utils

import (
	"math/rand"
	"time"
)

func RandString(length int) string {
	const charset = "1234567890"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
