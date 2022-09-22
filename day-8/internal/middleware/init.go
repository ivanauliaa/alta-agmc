package middleware

import (
	"day6-hexagonal/pkg/util/validator"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	LogMiddleware(e)

	e.Validator = &validator.CustomValidator{
		Validator: validator.NewValidator(),
	}
}
