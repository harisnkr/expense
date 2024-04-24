package models

// Card represents a credit/debit card that a user might have
type Card struct {
	ID         string                 `bson:"_id"`
	Name       string                 `bson:"name"`
	IssuerBank string                 `bson:"issuer_bank"`
	Network    string                 `bson:"network"`
	Miles      Miles                  `bson:"miles"`
	Image      string                 `bson:"image"`
	Other      map[string]interface{} `bson:"other"` // property bag
}

// Miles refer to miles related info that a Card can have
type Miles struct {
	BonusMultiplier    float64         `bson:"bonus_multiplier"`
	BonusMultiplierCap float64         `bson:"bonus_multiplier_cap"`
	LocalMultiplier    float32         `bson:"multiplier"`
	OverseasMultiplier float32         `bson:"overseas_multiplier"`
	MinimumSpend       float32         `bson:"minimum_spend"`
	SpendCategories    []SpendCategory `bson:"spend_categories"`
}

// SpendCategory is an int representation of the spend category
type SpendCategory int

const (
	Dining SpendCategory = iota
	Travel
	Gas
	Groceries
	Utilities // TODO: split into SP, mobile, etc.
	Foreign
	Contactless
	PublicTransport
	Online // TODO: split into food deliveries, ride hailing, etc.
	CharitableDonations
	Education
	Hospitals
	Hotels
	InsurancePremiums
	// TODO: are there more?
)

// String returns the string equivalent of the SpendCategory int
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
	case CharitableDonations:
		return "CharitableDonations"
	case Education:
		return "Education"
	case Hospitals:
		return "Hospitals"
	case Hotels:
		return "Hotels"
	case InsurancePremiums:
		return "InsurancePremiums"
	}

	return "Unknown"
}
