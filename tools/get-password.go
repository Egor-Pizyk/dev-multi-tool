package tools

import (
	"math/rand"
)

// TODO: Improve current realization (use mask?)
// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go

const LETTERS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GetPassword(size int) string {
	result := make([]byte, size)
	for i := 0; i <= size-1; i++ {
		result[i] = LETTERS[rand.Intn(size)]
	}
	return string(result)
}
