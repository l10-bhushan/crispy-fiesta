package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/l10-bhushan/crispy-fiesta/internal/models"
)

// Creating a context key type
type authContextKey string

// here userContextKey is of authContextKey type with value user
const userContextKey authContextKey = "user"

// Writing the AuthMiddleware
func AuthMiddleware(jwtSecret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// Fetching the AuthHeader value for "Authorization"
			authHeader := r.Header.Get("Authorization")

			// If we get nothing we throw an error
			if authHeader == "" {
				http.Error(w, "missing authorization header", http.StatusUnauthorized)
				return
			}

			// Splitting authHeader to separater Bearer and token
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(w, "invalid authorization header", http.StatusUnauthorized)
				return
			}

			// Storing the token into tokenString
			tokenString := parts[1]

			// Creating an instance of JWTClaims to store our user information
			claims := &models.JWTClaims{}

			// Parsing the token
			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				if token.Method != jwt.SigningMethodHS256 {
					return nil, jwt.ErrSignatureInvalid
				}

				return []byte(jwtSecret), nil
			})

			// Checking if the token is valid
			if err != nil || !token.Valid {
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}

			// Storing claims we go from ParseWithClaims into "user" using context
			ctx := context.WithValue(r.Context(), userContextKey, claims)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// Function to GetClaims from the context
func GetClaims(ctx context.Context) (*models.JWTClaims, bool) {

	// Fetch the claims from context
	claims, ok := ctx.Value(userContextKey).(*models.JWTClaims)

	return claims, ok
}
