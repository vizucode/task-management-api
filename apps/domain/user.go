package domain

// User represents user data for registration and login
type User struct {
	Username string `json:"username" example:"john_doe"`
	Email    string `json:"email" example:"john@example.com" validate:"required,email"`
	Password string `json:"password" example:"password123" validate:"required,min=6"`
}

// SignInResponse represents the response after successful login
type SignInResponse struct {
	AccessToken string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	Username    string `json:"username" example:"john_doe"`
}
