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

	staff := api.Group("/staff")
	{
		staff.POST("/create", controllers.CreateStaff) //entry point
		staff.POST("/update/:id", controllers.UpdateStaff)
		staff.POST("/delete/:id", controllers.DestroyStaff)
		staff.GET("/show/:id", controllers.GetStaff)
	}

}
