package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

// JWTCustomClaims model for custom claim
type JWTCustomClaims struct {
	Email string `json:"email"`
	ID    int32  `json:"id"`
	jwt.RegisteredClaims
}
