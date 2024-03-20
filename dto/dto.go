package dto

// CreateCardReq ...
type CreateCardReq struct {
	IssuerBank     string
	Name           string
	Network        string
	MilesPerDollar int
}

// UpdateCardReq ...
type UpdateCardReq struct {
	IssuerBank     string
	Name           string
	Network        string
	MilesPerDollar int
}
