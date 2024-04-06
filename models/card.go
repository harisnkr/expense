package models

type Card struct {
	Name       string `bson:"name"`
	IssuerBank string `bson:"issuer_bank,omitempty"`
	Network    string `bson:"network"`
	Miles      Miles  `bson:"miles"`
}

type Miles struct {
	Multiplier      float32         `bson:"multiplier"` // 2x, 1.4x, 4x json
	MinimumSpend    float32         `bson:"minimum_spend"`
	SpendCategories []SpendCategory `bson:"spend_categories"`
}

type SpendCategory int

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
