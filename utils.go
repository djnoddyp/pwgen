package main

import (
	"crypto/rand"
	"math/big"
	"fmt"
)

// Filter items based on a test
func Filter(data []string, pred func(s string) bool) (filtered []string) {
	for _, word := range data {
		if pred(word) {
			filtered = append(filtered, word)
		}
	}
	return
}

func RandomNumber(maxLen int) int64 {
	i, err := rand.Int(rand.Reader, big.NewInt(int64(maxLen)))
	if err != nil {
		panic(fmt.Sprintf("Cannot get random number, details: %s", err))
	}
	return i.Int64()
}
