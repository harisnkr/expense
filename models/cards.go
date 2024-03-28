package models

import "gorm.io/gorm"

type Card struct {
	gorm.Model
	IssuerBank     string  // Citibank
	Name           string  // Rewards, Preferred Platinum
	Network        string  // MasterCard World, Visa Platinum, Visa Signature
	MilesPerDollar float32 // 1.2, 4

}

type Miles struct {
	Multiplier      float32 // 2x, 1.4x, 4x
	MinimumSpend    float32
	SpendCategories []SpendCategory
}

type SpendCategory int64

const (
	Dining SpendCategory = iota
	Travel
	Gas
	Groceries
	Utilities
	Foreign
	Contactless
	PublicTransport
	Online
	// TODO: are there more?
)

func (s SpendCategory) String() string {
	switch s {
	case Dining:
		return "Dining"
	case Travel:
		return "Travel"
	case Gas:
		return "Gas"
	case Groceries:
		return "Groceries"
	case Utilities:
		return "Utilities"
	case Contactless:
		return "Contactless"
	case Foreign:
		return "Foreign"
	case PublicTransport:
		return "PublicTransport"
	case Online:
		return "Online"
	}
	return "Unknown"
}
