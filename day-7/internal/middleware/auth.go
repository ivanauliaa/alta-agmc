package middleware

import (
	"context"
	"net/http"
	"strings"

	"day6-hexagonal/internal/infrastructure/jwt"

	"github.com/labstack/echo/v4"
)

func Authentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authToken := c.Request().Header.Get("Authorization")
		if authToken == "" {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "unauthorized",
			})
		}

		token := strings.Split(authToken, "Bearer ")[1]
		id, err := jwt.ExtractPayload(token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "unauthorized",
			})
		}

		c.SetRequest(c.Request().WithContext(context.WithValue(c.Request().Context(), "userId", id)))

		return next(c)
	}
}
