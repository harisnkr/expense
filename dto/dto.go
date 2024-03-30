package dto

// RegisterUserRequest is the request body for /user/register
type RegisterUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=20,password"`
}
