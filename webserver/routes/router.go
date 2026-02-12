package routes

import (
	"encoding/json"
	"library/package/validator"
	"library/webserver/middlewares"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/tzdit/sample_api/package/util"
)

// Routers function
func Routers(app *echo.Echo) {

	//Common middleware
	app.Use(middlewares.Cors())
	app.Use(middlewares.Gzip())
	app.Use(middlewares.Logger(true))
	app.Use(middlewares.Secure())
	app.Use(middlewares.Recover())
	//app.Use(middlewares.KeyAuth())
	//app.Use(middlewares.Session())
	//app.Use(middlewares.AuditTrail())
	//app.Use(middlewares.TokenRefresherMiddleware())
	//app.Use(middlewares.JWTAuth())

	//initialize custom validator
	app.Validator = validator.GetValidator()

	//web routers
	ApiRouters(app)
	go generateRoutes(app)

}

func generateRoutes(e *echo.Echo) {
	basePath, _ := os.Getwd()
	dirPath := basePath + "/.storage/routes"
	filePath := dirPath + "/library.json"

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		os.MkdirAll(dirPath, 0755)
	}
	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		util.CheckError(err)
		return
	}
	err = os.WriteFile(filePath, data, 0644)
	util.CheckError(err)
}
