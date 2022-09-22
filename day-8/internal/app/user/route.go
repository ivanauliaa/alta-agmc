package user

import (
	"day6-hexagonal/internal/middleware"

	"github.com/labstack/echo/v4"
)

func (h *handler) Route(g *echo.Group) {
	g.POST("", h.Create)
	g.GET("", h.Get, middleware.Authentication)
	g.GET("/:id", h.GetById, middleware.Authentication)
	g.PUT("/:id", h.Update, middleware.Authentication)
	g.DELETE("/:id", h.Delete, middleware.Authentication)
}
