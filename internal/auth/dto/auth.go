package dto

// Request DTO

// RegisterRequest represents the registration request
// @Description Registration request
// @Accept json
// @Produce json
// @Param username body string true "Username"
// @Param email body string true "Email"
// @Param password body string true "Password"
type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// LoginRequest represents the login request
// @Description Login request
// @Accept json
// @Produce json
// @Param email body string true "Email"
// @Param password body string true "Password"
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse represents the login response
// @Description Login response
// @Accept json
// @Produce json
// @Param token body string true "Token"
type LoginResponse struct {
	Token string `json:"token"`
}

// AuthResponse represents the auth response
// @Description Auth response
// @Accept json
// @Produce json
// @Param message body string true "Message"
type AuthResponse struct {
	Message string `json:"message"`
}
