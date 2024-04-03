package common

import (
	"fmt"
	"math/rand"
)

// GenerateToken generates a random verification token
func GenerateToken() string {
	// Generate a random 8-digit number
	otp := rand.Intn(100000000)

	// Ensure the OTP is exactly 8 digits long
	return fmt.Sprintf("%08d", otp)
}
