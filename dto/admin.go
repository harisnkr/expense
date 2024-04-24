package dto

// AdminDeleteCardRequest is the request body to delete a card listed
type AdminDeleteCardRequest struct {
	Name       string `json:"name"`
	IssuerBank string `json:"issuerBank"`
	Network    string `json:"network"`
}

// AdminUpdateCardRequest is the request body to update a card
type AdminUpdateCardRequest struct {
	ID      string                 `binding:"required" json:"id"`
	Updates map[string]interface{} `binding:"required" json:"updates"`
}
