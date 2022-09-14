package routes

import (
	"day2-dynamic/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	v1Routes := e.Group("/v1")
	v1UserRoutes := v1Routes.Group("/users")

	v1UserRoutes.GET("", controllers.GetUsersController)
	v1UserRoutes.POST("", controllers.CreateUserController)
	v1UserRoutes.GET("/:id", controllers.GetUserByIdController)
	v1UserRoutes.PUT("/:id", controllers.UpdateUserByIdController)
	v1UserRoutes.DELETE("/:id", controllers.DeleteUserByIdController)

	return e
}
