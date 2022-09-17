package routes

import (
	"day3-middleware/controllers"
	"day3-middleware/middlewares"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(middlewares.GetJwtSecret()),
		Skipper: func(c echo.Context) bool {
			return skipped(&c)
		},
	}))

	v1Routes := e.Group("/v1")
	v1UserRoutes := v1Routes.Group("/users")

	v1UserRoutes.GET("", controllers.GetUsersController)
	v1UserRoutes.POST("", controllers.CreateUserController)
	v1UserRoutes.GET("/:id", controllers.GetUserByIdController)
	v1UserRoutes.PUT("/:id", controllers.UpdateUserByIdController)
	v1UserRoutes.DELETE("/:id", controllers.DeleteUserByIdController)

	v1BookRoutes := v1Routes.Group("/books")

	v1BookRoutes.GET("", controllers.GetBooksController)
	v1BookRoutes.POST("", controllers.CreateBookController)
	v1BookRoutes.GET("/:id", controllers.GetBookByIdController)
	v1BookRoutes.PUT("/:id", controllers.UpdateBookByIdController)
	v1BookRoutes.DELETE("/:id", controllers.DeleteBookByIdController)

	v1Routes.POST("/login", controllers.LoginController)

	middlewares.LogMiddleware(e)

	middlewares.AssignValidator(e)

	return e
}

func skipped(c *echo.Context) bool {
	return ((*c).Request().Method == "POST" && strings.Contains((*c).Request().URL.Path, "/v1/users")) ||
		((*c).Request().Method == "POST" && strings.Contains((*c).Request().URL.Path, "/v1/login")) ||
		((*c).Request().Method == "GET" && strings.Contains((*c).Request().URL.Path, "/v1/books"))
}
