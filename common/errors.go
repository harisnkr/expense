package common

import "errors"

var (
	// TODO: Add more error codes and implement them in the controllers
	EmailAlreadyTaken       = errors.New("email already taken")
	UnknownError            = errors.New("unknown error")
	InvalidVerificationCode = errors.New("invalid verification code")
)
