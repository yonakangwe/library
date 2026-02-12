package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

const (
	jwtSecretKey = "0ifp5bw(!1j8bq#2bwd24)bn!0$gco6hhoce^!7tmprdaf$1z7"
)

func GetJWTSecret() string {
	return jwtSecretKey
}
func JWTAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if SkipperJWTCheck(c) {
				return next(c)
			}
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error":   "missing Authorization header",
					"message": "missing Authorization header",
				})
			}
			if !strings.HasPrefix(authHeader, "Bearer ") {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error":   "invalid Authorization header format",
					"message": "invalid Authorization header format",
				})
			}
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("unexpected signing method")
				}
				return []byte(GetJWTSecret()), nil
			})

			if err != nil || !token.Valid {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error":   "invalid or expired token",
					"message": "invalid or expired token",
				})
			}
			return next(c)
		}
	}
}

func SkipperJWTCheck(c echo.Context) bool {
	if strings.HasSuffix(c.Path(), "/login") ||
		strings.HasSuffix(c.Path(), "/register") ||
		strings.HasSuffix(c.Path(), "/reset-password-token") ||
		strings.HasSuffix(c.Path(), "/verify-email") ||
		strings.HasSuffix(c.Path(), "/re-fresh-token") ||
		strings.HasSuffix(c.Path(), "/l-gout") ||
		strings.HasSuffix(c.Path(), "/auth-subsystem-get-user") ||
		strings.HasSuffix(c.Path(), "/request-reset-password-link") {
		return true
	}
	return false
}
