package tools

import (
	"crypto/rand"
	"math/big"
)

// TODO: Improve current realization (use mask?)
// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go

const LETTERS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const NUMBERS = "0123456789"

func GetPassword(size int, isNumbersInUsage bool) (string, error) {
	charset := LETTERS
	if isNumbersInUsage {
		charset += NUMBERS
	}

	result := make([]byte, size)
	for i := 0; i <= size-1; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result[i] = charset[n.Int64()]
	}
	return string(result), nil
}
