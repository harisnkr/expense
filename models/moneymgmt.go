package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Budget struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    primitive.ObjectID `bson:"user_id"`
	Category  string             `bson:"category"`
	Amount    float64            `bson:"amount"`
	StartDate time.Time          `bson:"start_date"`
	EndDate   time.Time          `bson:"end_date"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

type Transaction struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	UserID      primitive.ObjectID `bson:"user_id"`
	Type        string             `bson:"type"`
	Category    string             `bson:"category"`
	Amount      float64            `bson:"amount"`
	Date        time.Time          `bson:"date"`
	Description string             `bson:"description"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}

type Savings struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	UserID        primitive.ObjectID `bson:"user_id"`
	Description   string             `bson:"description"`
	TargetAmount  float64            `bson:"target_amount"`
	CurrentAmount float64            `bson:"current_amount"`
	Deadline      time.Time          `bson:"deadline"`
	CreatedAt     time.Time          `bson:"created_at"`
	UpdatedAt     time.Time          `bson:"updated_at"`
}
