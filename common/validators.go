package common

import (
	"net/mail"
	"regexp"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// InitValidators ...
func InitValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation(Password, validatePassword)
		_ = v.RegisterValidation(Username, validateUsername)
		_ = v.RegisterValidation(Email, validateEmail)
	}
}

func validatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	if len(password) < 8 || len(password) > 20 {
		return false
	}
	// TODO: more password validation logic here
	// For example, checking for special characters, uppercase, lowercase, etc.
	return true
}

func validateUsername(fl validator.FieldLevel) bool {
	username := fl.Field().String()

	// Define your regex pattern for username validation
	regex := regexp.MustCompile(`^[a-zA-Z0-9_-]{4,16}$`)
	// TOOD: blacklist words
	return regex.MatchString(username)
}

// validateEmail ...
func validateEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String()

	// Define your regex pattern for email validation
	_, err := mail.ParseAddress(email)
	return err == nil
}
