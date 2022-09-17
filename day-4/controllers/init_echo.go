package controllers

import (
	"day3-middleware/config"
	"day3-middleware/middlewares"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitEcho() *echo.Echo {
	config.InitDB()

	e := echo.New()

	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(middlewares.GetJwtSecret()),
		Skipper: func(c echo.Context) bool {
			return skipped(&c)
		},
	}))

	middlewares.AssignValidator(e)

	return e
}

func skipped(c *echo.Context) bool {
	return ((*c).Request().Method == "POST" && strings.Contains((*c).Request().URL.Path, "/v1/users")) ||
		((*c).Request().Method == "POST" && strings.Contains((*c).Request().URL.Path, "/v1/login")) ||
		((*c).Request().Method == "GET" && strings.Contains((*c).Request().URL.Path, "/v1/books"))
}
