package main

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type UserResponse struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	IsActive int    `json:"isActive"`
}
