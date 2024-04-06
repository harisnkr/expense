package models

import (
	"time"
)

type User struct {
	ID        string    `bson:"_id"`
	FirstName string    `bson:"first_name"`
	LastName  string    `bson:"last_name"`
	Username  string    `bson:"username"`
	Password  string    `bson:"password"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`

	Email              string    `bson:"email"`
	Verified           bool      `bson:"verified"`
	VerificationCode   string    `bson:"verification_code"`
	VerificationSentAt time.Time `bson:"verification_sent_at"`

	Cards        []Card        `bson:"cards"`
	Budgets      []Budget      `bson:"budgets"`
	Transactions []Transaction `bson:"transactions"`
	Savings      []Savings     `bson:"savings"`
}
