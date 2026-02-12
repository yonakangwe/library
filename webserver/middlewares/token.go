package middlewares

import (
	"library/package/log"
	"library/package/wrappers"
	"library/webserver/auth"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

const (
	refreshTokenCookieName = "refresh-token"
	expirationTime         = 60
)

func TokenRefresherMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if SkipperLoginCheck(c) {
				return next(c)
			}
			if c.Get("user") == nil {
				return next(c)
			}
			u := c.Get("user").(*jwt.Token)
			claims := u.Claims.(*auth.JWTCustomClaims)

			if time.Until(time.Unix(claims.ExpiresAt.Unix(), 0)) < expirationTime/2*time.Minute {
				// Gets the refresh token from the cookie.
				rc, err := c.Cookie(refreshTokenCookieName)
				if err == nil && rc != nil {
					// Parses token and checks if it valid.
					tkn, err := jwt.ParseWithClaims(rc.Value, claims, func(token *jwt.Token) (interface{}, error) {
						return []byte(auth.GetRefreshJWTSecret()), nil
					})
					if err != nil {
						if err == jwt.ErrSignatureInvalid {
							return wrappers.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
						}
					}
					if tkn != nil && tkn.Valid {
						_, _, _, err := auth.GenerateTokensAndSetCookies(claims.ID, claims.Email, c)
						if err != nil {
							log.Errorf("error generating token: %v", err)
							return wrappers.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
						}
					}
				}
			}
			return next(c)
		}
	}
}

func SkipperLoginCheck(c echo.Context) bool {
	if strings.HasSuffix(c.Path(), "/login") ||
		strings.HasSuffix(c.Path(), "/register") ||
		strings.HasSuffix(c.Path(), "/reset-password-token") ||
		strings.HasSuffix(c.Path(), "/verify-email") ||
		strings.HasSuffix(c.Path(), "/request-reset-password-link") {
		return true
	}
	return false
}
