package routes

import (
	"day2-static/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	v1Routes := e.Group("/v1")
	v1BookRoutes := v1Routes.Group("/books")

	v1BookRoutes.GET("", controllers.GetBooksController)
	v1BookRoutes.POST("", controllers.CreateBookController)
	v1BookRoutes.GET("/:id", controllers.GetBookByIdController)
	v1BookRoutes.PUT("/:id", controllers.UpdateBookByIdController)
	v1BookRoutes.DELETE("/:id", controllers.DeleteBookByIdController)

	return e
}
