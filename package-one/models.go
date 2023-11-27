package package_one

import (
	"math/rand"
)

const Version = "v0.0"


type DataStructTypeOne struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

func GenerateRandomData() DataStructTypeOne {
	return DataStructTypeOne{
		ID:      randSeq(8),
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


