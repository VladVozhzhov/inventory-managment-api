package utils

import (
	"math/rand"
)

const idCharset = "abcdefghijklmnopqrstuvwxyz0123456789"
const idLength = 20

func GenerateRandomID() string {
	b := make([]byte, idLength)
	for i := range b {
		b[i] = idCharset[rand.Intn(len(idCharset))]
	}
	return string(b)
}
