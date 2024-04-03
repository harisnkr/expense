package models

import (
	"time"
)

type User struct {
	ID        string    `bson:"_id"`
	FirstName string    `json:"first_name,omitempty" bson:"first_name"`
	LastName  string    `json:"last_name,omitempty" bson:"last_name"`
	Username  string    `json:"username" bson:"username" binding:"required,username"`
	Password  string    `json:"password" bson:"password" binding:"required,password"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

	Email              string    `json:"email" binding:"required,email"`
	Verified           bool      `json:"verified"`
	VerificationCode   string    `json:"verification_code" bson:"verification_code"`
	VerificationSentAt time.Time `json:"verification_sent_at" bson:"verification_sent_at"`

	Cards        []Card        `bson:"cards" json:"cards"`
	Budgets      []Budget      `json:"budgets,omitempty" bson:"budgets"`
	Transactions []Transaction `json:"transactions,omitempty" bson:"transactions"`
	Savings      []Savings     `json:"savings" bson:"savings"`
}
