package webserver

import (
	"fmt"
	"library/config"
	"library/package/helpers"
	"library/package/log"
	"library/webserver/middlewares"
	"library/webserver/routes"
	"library/webserver/systems"

	"github.com/labstack/echo/v4"
)

// StartMkoaServer starts a server that only serves the mkoa page and API
func StartMkoaServer() {
	e := echo.New()
	e.HideBanner = true

	// Minimal middlewares used by the project
	e.Use(middlewares.Cors())
	e.Use(middlewares.Gzip())
	e.Use(middlewares.Logger(true))
	e.Use(middlewares.Secure())
	e.Use(middlewares.Recover())

	// Routes: serve mkoa page and mkoa API group
	routes.MkoaPage(e)
	api := e.Group("library/api/v1")
	routes.MkoaRoutes(api)

	// Init helpers and systems
	helpers.Init()
	systems.Init()

	e.Debug = true

	cfg, err := config.New()
	if err != nil {
		log.Errorf("error getting config: %v", err)
	}
	address := fmt.Sprintf("%v:%v", cfg.WebServer.PublicHost, cfg.WebServer.Port)
	e.Logger.Fatal(e.Start(address))
}
