package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func Token() {
	secret := make([]byte, 64)
	_, err := rand.Read(secret)
	if err != nil {
		panic(err)
	}
	fmt.Println(hex.EncodeToString(secret))
}
