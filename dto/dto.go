package dto

// RegisterUserRequest is the request body for /user/register.
type RegisterUserRequest struct {
	FirstName string `binding:"required,name"     json:"firstName"`
	LastName  string `binding:"required,name"     json:"lastName"`
	Email     string `binding:"required,email"    json:"email"`
	Password  string `binding:"required,password" json:"password"`
}

// UserEmailVerifyRequest is the request body for POST /user/email/verify
type UserEmailVerifyRequest struct {
	Email            string `binding:"required,email" json:"email"`
	VerificationCode string `binding:"required"       json:"verificationCode"`
}

// UserLoginResponse is the response body /user/login
type UserLoginResponse struct {
	SessionToken string `json:"sessionToken"`
	ExpiresIn    string `json:"expiresIn"`
}

// AdminDeleteCardRequest is the request body to delete a card listed
type AdminDeleteCardRequest struct {
	Name       string `json:"name"`
	IssuerBank string `json:"issuerBank"`
	Network    string `json:"network"`
}

// GetCardRequest ...
type GetCardRequest struct {
	Name string `json:"name"`
}
