package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Cors() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://192.168.26.18", // auth
			"http://192.168.26.10", // repo
		},
		AllowMethods: []string{echo.OPTIONS, echo.POST, echo.GET},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			"SYSTEM-NAME", "DATA-HASH", "DATA-SIGNATURE", "Authorization", "SYSTEM-CLIENT-ID",
		},
		AllowCredentials: true, // needed if you send cookies or auth headers
		MaxAge:           86400,
	})
}
