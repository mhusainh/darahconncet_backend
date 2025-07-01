package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

// GenerateCertificateNumber generates a unique 20-digit certificate number
// Format: YYYYMMDDHHMMSS + 6 random digits
func GenerateCertificateNumber() (string, error) {
	// Get current timestamp (14 digits: YYYYMMDDHHMMSS)
	now := time.Now()
	timestamp := now.Format("20060102150405")

	// Generate 6 random digits
	randomPart, err := generateRandomDigits(6)
	if err != nil {
		return "", fmt.Errorf("failed to generate random digits: %w", err)
	}

	// Combine timestamp + random digits = 20 digits total
	certificateNumber := timestamp + randomPart

	return certificateNumber, nil
}

// generateRandomDigits generates n random digits as a string
func generateRandomDigits(n int) (string, error) {
	if n <= 0 {
		return "", fmt.Errorf("number of digits must be positive")
	}

	// Calculate the maximum value for n digits (10^n - 1)
	max := new(big.Int)
	max.Exp(big.NewInt(10), big.NewInt(int64(n)), nil)

	// Generate random number between 0 and max-1
	randomNum, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", err
	}

	// Format with leading zeros to ensure exactly n digits
	format := fmt.Sprintf("%%0%dd", n)
	return fmt.Sprintf(format, randomNum), nil
}

// GenerateUniqueCertificateNumber generates a certificate number with additional uniqueness
// This version includes nanoseconds for even higher uniqueness
func GenerateUniqueCertificateNumber() (string, error) {
	// Get current timestamp with nanoseconds
	now := time.Now()
	
	// Use YYYYMMDDHHMMSS (14 digits)
	timestamp := now.Format("20060102150405")
	
	// Add last 2 digits of nanoseconds for extra uniqueness
	nanos := now.Nanosecond() % 100
	nanosStr := fmt.Sprintf("%02d", nanos)
	
	// Generate 4 random digits
	randomPart, err := generateRandomDigits(4)
	if err != nil {
		return "", fmt.Errorf("failed to generate random digits: %w", err)
	}

	// Combine: timestamp(14) + nanos(2) + random(4) = 20 digits total
	certificateNumber := timestamp + nanosStr + randomPart

	return certificateNumber, nil
}