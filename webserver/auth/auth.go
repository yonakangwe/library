package auth

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

const (
	jwtSecretKey        = "0ifp5bw(!1j8bq#2bwd24)bn!0$gco6hhoce^!7tmprdaf$1z7"
	jwtRefreshSecretKey = "0ifp5bw(!1j7bq#2bwd24)bn!0$gco5hhoce^!7tmprdaf$1z7"
)

func GetJWTSecret() string {
	return jwtSecretKey
}

func GetRefreshJWTSecret() string {
	return jwtRefreshSecretKey
}

func GetUserFromContext(c echo.Context) (userID int32, emailAddress string) {
	var tokenString string
	authHeader := c.Request().Header.Get("Authorization")

	if after, ok := strings.CutPrefix(authHeader, "Bearer "); ok {
		tokenString = after
	} else {
		c.Response().WriteHeader(http.StatusUnauthorized)
		return userID, emailAddress
	}

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			c.Response().WriteHeader(http.StatusUnauthorized)
		}
		return []byte(GetJWTSecret()), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return int32(claims["id"].(float64)), claims["email"].(string)
	}
	return userID, emailAddress
}
