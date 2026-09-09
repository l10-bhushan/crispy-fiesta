package models

import (
	"time"

	"github.com/google/uuid"
)

// Request structure for finding the user by email
type FindByEmailRequest struct {
	Email string `json:"email"`
}

// Request structure for registering a user
type RegisterUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// Response structure for user
type UserResponse struct {
	Id        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

// Request struct for user login
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
