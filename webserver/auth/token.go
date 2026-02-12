package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

const (
	accessTokenExpiryMinutes  = 15 // 15min
	refreshTokenExpiryMinutes = 60 // 60min
)

func generateToken(id int32, email string, expTime time.Time, secret []byte) (string, time.Time, error) {
	claims := &JWTCustomClaims{
		ID:    id,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", time.Time{}, err
	}

	return tokenString, expTime, nil
}
func GenerateAccessToken(id int32, email string) (string, time.Time, error) {
	expTime := time.Now().Add(accessTokenExpiryMinutes * time.Minute)
	return generateToken(id, email, expTime, []byte(GetJWTSecret()))
}

func GenerateRefreshToken(id int32, email string) (string, time.Time, error) {
	expTime := time.Now().Add(refreshTokenExpiryMinutes * time.Minute)
	return generateToken(id, email, expTime, []byte(GetRefreshJWTSecret()))
}

func GenerateTokensAndSetCookies(id int32, email string, c echo.Context) (string, string, time.Time, error) {
	accessToken, accessExp, err := GenerateAccessToken(id, email)
	if err != nil {
		return "", "", time.Time{}, err
	}
	setTokenCookie(email+accessTokenCookieName, accessToken, accessExp, c)
	setUserCookie(email, accessExp, c)

	refreshToken, refreshExp, err := GenerateRefreshToken(id, email)
	if err != nil {
		return "", "", time.Time{}, err
	}
	setTokenCookie(email+refreshTokenCookieName, refreshToken, refreshExp, c)
	return accessToken, refreshToken, refreshExp, nil
}
