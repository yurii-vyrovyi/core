package package_two

import (
	"math/rand"
	"time"
)

const Version = "v0.2"

type DataStructTypeOne struct {
	ID      time.Time `json:"id"`
	Message string `json:"message"`
}

func GenerateRandomData() DataStructTypeOne {
	return DataStructTypeOne{
		ID:      time.Now(),
		Message: randSeq(12),
	}
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
