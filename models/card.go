package models

type Card struct {
	Name       string `json:"name"`
	IssuerBank string `json:"issuer_bank,omitempty"`
	Network    string `json:"network"`
	Miles      Miles  `json:"miles" json:"miles"`
}

type Miles struct {
	Multiplier      float32         `json:"multiplier"` // 2x, 1.4x, 4x json
	MinimumSpend    float32         `json:"minimumSpend"`
	SpendCategories []SpendCategory `json:"spendCategories" bson:"spendCategories"`
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
