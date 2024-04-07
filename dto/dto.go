package dto

// RegisterUserRequest is the request body for /user/register.
type RegisterUserRequest struct {
	FirstName string `binding:"required,name"     json:"firstName"`
	LastName  string `binding:"required,name"     json:"lastName"`
	Email     string `binding:"required,email"    json:"email"`
	Password  string `binding:"required,password" json:"password"`
}

// AdminCreateCardRequest is the request body for POST /admin/card.
type AdminCreateCardRequest struct {
	Name         string  `json:"name"`
	IssuerBank   string  `json:"issuerBank"`
	Network      string  `json:"network"`
	Multiplier   float32 `json:"multiplier"`
	MinimumSpend float32 `json:"minimumSpend"`
	SpendLimit   []int   `json:"spendLimit"`
}

// UserLoginRequest is the request body for POST /user/login
type UserLoginRequest struct {
	Email    string `binding:"required,email"    json:"email"`
	Password string `binding:"required,password" json:"password"`
}

type UserLoginResponse struct {
	SessionToken string `json:"sessionToken"`
	ExpiresIn    string `json:"expiresIn"`
}
