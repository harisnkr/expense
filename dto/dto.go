package dto

// RegisterUserRequest is the request body for /user/register
type RegisterUserRequest struct {
	FirstName string `json:"firstName,omitempty" binding:"required,name"`
	LastName  string `json:"lastName,omitempty" binding:"required,name"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=8,max=20,password"`
}
