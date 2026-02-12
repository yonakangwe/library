package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Cors() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			// "http://102.208.184.52",                  // frontend server IP
			// "http://102.208.184.51",                  // frontend server IP
			// "http://localhost:5000",                  // local dev
			//"http://localhost:5001", // another dev frontend
			//"http://localhost:5002", // another dev frontend
			// "http://localhost:4000",                  // local dev
			// "http://localhost:4001",                  // another dev frontend
			"https://nissti.costech.or.tz",     // another production frontend
			"https://backoffice.costech.or.tz", // another production frontend
			"http://192.168.26.18",             // auth
			"http://192.168.26.10",             // repo

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
