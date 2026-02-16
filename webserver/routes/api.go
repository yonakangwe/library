package routes

import (
	"library/webserver/controllers"

	"github.com/labstack/echo/v4"
)

func ApiRouters(app *echo.Echo) {

	/*
		|--------------------------------------------------------------------------
		| API Routes
		|--------------------------------------------------------------------------
	*/
	api := app.Group("library/api/v1")

	/*
		|--------------------------------------------------------------------------
		| DONE - role
		|--------------------------------------------------------------------------
	*/
	role := api.Group("/role")
	{
		role.POST("/list", controllers.ListRole)
		role.POST("/show", controllers.GetRole)
		role.POST("/create", controllers.CreateRole)
		role.POST("/update", controllers.UpdateRole)
		role.POST("/delete", controllers.SoftDeleteRole)
		role.POST("/destroy", controllers.DestroyRole)
	}

	member := api.Group("/member")
	{
		member.POST("/create", controllers.CreateMember)
	}

	/*
		|--------------------------------------------------------------------------
		| DONE - country
		|--------------------------------------------------------------------------
	*/
	countryRoute := api.Group("/country")
	{
		countryRoute.POST("/list", controllers.ListCountry)
		countryRoute.POST("/show", controllers.GetCountry)
		countryRoute.POST("/create", controllers.CreateCountry)
		countryRoute.POST("/update", controllers.UpdateCountry)
		countryRoute.POST("/delete", controllers.SoftDeleteCountry)
		countryRoute.POST("/destroy", controllers.DestroyCountry)
	}

}
