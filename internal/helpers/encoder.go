package helpers

import (
	"math/rand"
	"time"
)

var (
	CharacterSet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz_"
)

func Encode(num int) string {
	b := make([]byte, num)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = CharacterSet[rand.Intn(len(CharacterSet))]
	}
	return string(b)
}
