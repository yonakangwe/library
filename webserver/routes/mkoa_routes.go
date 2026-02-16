package routes

import (
	"library/webserver/controllers"

	"github.com/labstack/echo/v4"
)

/*
|--------------------------------------------------------------------------
| DONE - mkoa (API)
|--------------------------------------------------------------------------
*/
func MkoaRoutes(api *echo.Group) {
	mkoa := api.Group("/mkoa")
	{
		mkoa.POST("/create", controllers.CreateMkoa)
		mkoa.POST("/list", controllers.ListMkoa)
		mkoa.POST("/show", controllers.GetMkoa)
		mkoa.POST("/update", controllers.UpdateMkoa)
		mkoa.POST("/delete", controllers.SoftDeleteMkoa)
		mkoa.POST("/destroy", controllers.DestroyMkoa)
	}
}

/*
|--------------------------------------------------------------------------
| DONE - mkoa (browser page)
|--------------------------------------------------------------------------
| Serves the mkoa list page at GET /mkoa. Run the app from project root so
| the path "webserver/public/mkoa/index.html" resolves correctly.
*/
func MkoaPage(app *echo.Echo) {
	app.GET("/", func(c echo.Context) error { return c.Redirect(302, "/mkoa") })
	app.Static("/css", "webserver/public/css")
	app.File("/mkoa", "webserver/public/mkoa/index.html")
}
