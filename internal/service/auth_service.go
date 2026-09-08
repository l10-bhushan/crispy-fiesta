package service

import (
	"time"

	"github.com/google/uuid"

	"github.com/golang-jwt/jwt/v5"
	"github.com/l10-bhushan/crispy-fiesta/internal/models"
)

// AuthService struct that will hold our jwtSecret
type AuthService struct {
	jwtSecret []byte
}

// Contructor function for AuthService struct
func NewAuthService(jwtSecret string) *AuthService {
	return &AuthService{
		jwtSecret: []byte(jwtSecret),
	}
}

// Service to generate JWT token
func (a *AuthService) GenerateToken(userId uuid.UUID) (string, error) {

	// Creating the claims section of the JWT token
	claims := models.JWTClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Adding the signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Signing the generated token with our jwt Secret
	return token.SignedString(a.jwtSecret)
}
