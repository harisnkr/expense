package common

import (
	"net/mail"
	"regexp"

	"github.com/go-playground/validator/v10"
)

// ValidatePassword ...
func ValidatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	if len(password) < 8 || len(password) > 20 {
		return false
	}

	// You can add more password validation logic here
	// For example, checking for special characters, uppercase, lowercase, etc.
	return true
}

// ValidateUsername ...
func ValidateUsername(fl validator.FieldLevel) bool {
	username := fl.Field().String()

	// Define your regex pattern for username validation
	regex := regexp.MustCompile(`^[a-zA-Z0-9_-]{4,16}$`)
	return regex.MatchString(username)
}

// ValidateEmail ...
func ValidateEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String()

	// Define your regex pattern for email validation
	_, err := mail.ParseAddress(email)
	return err == nil
}
