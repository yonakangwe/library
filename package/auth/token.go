package auth

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/golang-jwt/jwt"
)

// TokenBlacklist keeps track of invalidated tokens
var TokenBlacklist []string

// Custom claims structure to include email address
type CustomClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// secret key for signing the tokens
var secretKey = []byte("0ifp5bw(!1j8bq#2bwd24)bn!0$gco6hhoce^!7tmprdaf$1z7")

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// GenerateToken generates a JWT token with an expiration date and email address
func GenerateToken(email string, expireTime time.Time) (string, error) {
	// Create the claims
	claims := CustomClaims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}

	// Create the token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Sign the token with the secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ValidateToken validates if a JWT token has expired or not, and retrieves the email address
func ValidateToken(tokenString string) (string, error) {

	// Check if the token has been invalidated
	if IsTokenInvalid(tokenString) {
		return "", errors.New("token has been invalidated")
	}
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		return "", err
	}

	// Check if the token is valid and hasn't expired
	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	return claims.Email, nil
}

// InvalidateToken invalidates a token by adding it to the blacklist
func InvalidateToken(tokenString string) {
	TokenBlacklist = append(TokenBlacklist, tokenString)
}

// IsTokenInvalid checks if a token has been invalidated
func IsTokenInvalid(tokenString string) bool {
	for _, invalidatedToken := range TokenBlacklist {
		if invalidatedToken == tokenString {
			return true
		}
	}
	return false
}

func GenerateSecureRandomString(length int) (string, error) {
	result := make([]byte, length)
	charsetLength := big.NewInt(int64(len(charset)))

	for i := range result {
		num, err := rand.Int(rand.Reader, charsetLength)
		if err != nil {
			return "", err
		}
		result[i] = charset[num.Int64()]
	}

	return string(result), nil
}
