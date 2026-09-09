package models

import (
	"github.com/google/uuid"

	"github.com/golang-jwt/jwt/v5"
)

// Creating the claims struct, we can add other fields as well like name, email before jwt.Registerd claims
type JWTClaims struct {
	UserId uuid.UUID `json:"user_id"`
	Role   string    `json:"role"`
	jwt.RegisteredClaims
}
