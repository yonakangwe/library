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
		role.POST("/create", controllers.CreateRole)
		// role.POST("/list", controllers.ListRole)
		// role.POST("/show", controllers.GetRole)
		// role.POST("/update", controllers.UpdateRole)
		// role.POST("/delete", controllers.SoftDeleteRole)
		// role.POST("/destroy", controllers.DestroyRole)
	}

	/*
		|--------------------------------------------------------------------------
		| DONE - mkoa
		|--------------------------------------------------------------------------
	*/
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
