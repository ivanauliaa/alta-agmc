package http

import (
	"day6-hexagonal/internal/app/auth"
	"day6-hexagonal/internal/app/book"
	"day6-hexagonal/internal/app/user"
	"day6-hexagonal/internal/factory"

	"github.com/labstack/echo/v4"
)

func NewHttp(e *echo.Echo, f *factory.Factory) {
	v1G := e.Group("/v1")

	user.NewHandler(f).Route(v1G.Group("/users"))
	auth.NewHandler(f).Route(v1G.Group("/auth"))
	book.NewHandler(f).Route(v1G.Group("/books"))
}
