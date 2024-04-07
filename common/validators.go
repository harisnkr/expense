package common

import (
	"net/mail"
	"regexp"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func initValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation(Name, validateName)
		_ = v.RegisterValidation(Username, validateUsername)
		_ = v.RegisterValidation(Password, validatePassword)
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
	regex := regexp.MustCompile(`^[a-zA-Z0-9_-]{4,16}$`)
	// TOOD: blacklist words
	return regex.MatchString(username)
}

func validateEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	_, err := mail.ParseAddress(email)
	return err == nil
}

func validateName(fl validator.FieldLevel) bool {
	name := fl.Field().String()
	return !(len(name) > 20)
}
