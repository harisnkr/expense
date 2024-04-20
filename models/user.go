package models

import (
	"time"
)

// User represents a user onboarded or undergoing onboarding
type User struct {
	ID             string    `bson:"_id"`
	FirstName      string    `bson:"first_name"`
	LastName       string    `bson:"last_name"`
	Password       string    `bson:"password"`
	CreatedAt      time.Time `bson:"created_at"`
	UpdatedAt      time.Time `bson:"updated_at"`
	ProfilePicture string    `bson:"profile_picture"`

	Email              string    `bson:"email"`
	Verified           bool      `bson:"verified"`
	VerificationCode   string    `bson:"verification_code"`
	VerificationSentAt time.Time `bson:"verification_sent_at"`

	Cards        []Card        `bson:"cards"`
	Budgets      []Budget      `bson:"budgets"`
	Transactions []Transaction `bson:"transactions"`
	Savings      []Savings     `bson:"savings"`
}
