package auth

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

const (
	accessTokenCookieName  = "access-token"
	refreshTokenCookieName = "refresh-token"
	AuthContextKey         = "auth_user"
)

func setTokenCookie(name, token string, expiration time.Time, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = token
	cookie.Expires = expiration
	cookie.Path = "/"
	cookie.HttpOnly = true
	c.SetCookie(cookie)
}
func clearTokenCookie(name string, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = ""
	cookie.MaxAge = -1
	cookie.Path = "/"
	c.SetCookie(cookie)
}

func setUserCookie(email string, expiration time.Time, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = "user"
	cookie.Value = email
	cookie.Expires = expiration
	cookie.Path = "/"
	c.SetCookie(cookie)
}

func clearUserCookie(c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = "user"
	cookie.Value = ""
	cookie.MaxAge = -1
	cookie.Path = "/"
	c.SetCookie(cookie)
}

func ClearSession(c echo.Context) {
	_, emailAddress := GetUserFromContext(c)
	clearUserCookie(c)
	clearTokenCookie(emailAddress+accessTokenCookieName, c)
	clearTokenCookie(emailAddress+refreshTokenCookieName, c)
}
