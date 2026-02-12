package webserver

import (
	"library/config"
	"library/package/helpers"
	"library/package/log"
	"library/webserver/routes"
	"library/webserver/systems"
	"fmt"

	"github.com/labstack/echo/v4"
)

// StartWebserver starts a webserver
func StartWebserver() {
	// Echo instance
	e := echo.New()

	//Disable echo banner
	e.HideBanner = true

	// Routes
	routes.Routers(e)

	//Init cache
	helpers.Init()
	systems.Init()

	//Remove this in production
	e.Debug = true

	cfg, err := config.New()
	if err != nil {
		log.Errorf("error getting config: %v", err)
	}
	address := fmt.Sprintf("%v:%v", cfg.WebServer.PublicHost, cfg.WebServer.Port)
	e.Logger.Fatal(e.Start(address))
}
