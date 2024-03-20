package dto

// CreateCardReq ...
type CreateCardReq struct {
	IssuerBank     string `json:"issuerBank"`
	Name           string `json:"name"`
	Network        string `json:"network"`
	MilesPerDollar int    `json:"milesPerDollar"`
}

// UpdateCardReq ...
type UpdateCardReq struct {
	IssuerBank     string `json:"issuerBank"`
	Name           string `json:"name"`
	Network        string `json:"network"`
	MilesPerDollar int    `json:"milesPerDollar"`
}
