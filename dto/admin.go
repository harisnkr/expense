package dto

// AdminCreateCardRequest is the request body for POST /admin/card.
type AdminCreateCardRequest struct {
	Name         string  `json:"name"`
	IssuerBank   string  `json:"issuerBank"`
	Network      string  `json:"network"`
	Multiplier   float32 `json:"multiplier"`
	MinimumSpend float32 `json:"minimumSpend"`
	SpendLimit   []int   `json:"spendLimit"`
}
