package models

import "time"

// Request structure for finding the user by email
type FindByEmailRequest struct {
	Email string `json:"email"`
}

// Request structure for registering a user
type RegisterUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// Response structure for user
type UserResponse struct {
	Id        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
