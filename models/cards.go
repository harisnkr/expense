package models

import "gorm.io/gorm"

type Card struct {
	gorm.Model
	IssuerBank     string // Citibank
	Name           string // Rewards, Preferred Platinum
	Network        string // MasterCard World, Visa Platinum, Visa Signature
	MilesPerDollar int    // 1.2, 4
}
