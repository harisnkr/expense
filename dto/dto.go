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

// UserLoginRequest is the request body for POST /user/login
type UserLoginRequest struct {
	Email    string `binding:"required,email"    json:"email"`
	Password string `binding:"required,password" json:"password"`
}

// UserLoginResponse is the response body for POST /user/login
type UserLoginResponse struct {
	SessionToken string `json:"sessionToken"`
	ExpiresIn    string `json:"expiresIn"`
}

// AddCardToUserRequest is the request body for POST /user/card
type AddCardToUserRequest struct {
	CardID string `json:"cardID"`
}

// UpdateMeRequest is the request body for PATCH /user/profile
type UpdateMeRequest struct {
	FirstName      *string `binding:"required,name"           json:"firstName"`
	LastName       *string `binding:"required,name"           json:"lastName"`
	ProfilePicture *string `binding:"required,profilePicture" json:"profilePicture"`
}
