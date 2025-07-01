package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateRandomPassportNumber() string {
	rand.Seed(time.Now().UnixNano())
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"

	// Generate two random letters
	letter1 := letters[rand.Intn(len(letters))]
	letter2 := letters[rand.Intn(len(letters))]

	// Generate seven random digits
	digits := make([]byte, 7)
	for i := 0; i < 7; i++ {
		digits[i] = numbers[rand.Intn(len(numbers))]
	}

	return fmt.Sprintf("%c%c%s", letter1, letter2, string(digits))
}