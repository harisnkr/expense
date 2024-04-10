package common

import "errors"

var (
	EmailAlreadyTaken       = errors.New("email already taken")
	UnknownError            = errors.New("unknown error")
	InvalidVerificationCode = errors.New("invalid verification code")
)
