package middlewares

import (
	"library/config"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Logger Middleware
func Logger(debug bool) echo.MiddlewareFunc {
	path := config.LoggerPath()
	out, err := os.Create(path + "/requests.log")
	if err != nil || debug {
		out = os.Stdout
	}

	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "Method=${method}, Url=\"${uri}\", Status=${status}, Time=${time_custom}, Latency=${latency_human} \n",
		CustomTimeFormat: "2006-01-02T15:04:05",
		Output:           out,
	})
}
