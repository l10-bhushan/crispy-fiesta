package models

import "github.com/golang-jwt/jwt/v5"

// Creating the claims struct, we can add other fields as well like name, email before jwt.Registerd claims
type JWTClaims struct {
	UserId string `json:"user_id"`
	jwt.RegisteredClaims
}
