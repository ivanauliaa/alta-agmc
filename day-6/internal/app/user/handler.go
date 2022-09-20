package user

import (
	"day6-hexagonal/internal/dto"
	"day6-hexagonal/internal/factory"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type handler struct {
	service ServiceInterface
}

func NewHandler(f *factory.Factory) *handler {
	return &handler{
		service: NewService(f),
	}
}

func (h *handler) Create(c echo.Context) error {
	payload := new(dto.CreateUserRequest)

	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	if err := c.Validate(payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	result, err := h.service.Create(c.Request().Context(), payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status": result,
	})
}

func (h *handler) Get(c echo.Context) error {
	users, err := h.service.FindAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

func (h *handler) GetById(c echo.Context) error {
	payload := new(dto.ByIdRequest)
	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	if err := c.Validate(payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	user, err := h.service.FindById(c.Request().Context(), payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user",
		"user":    user,
	})
}

func (h *handler) Update(c echo.Context) error {
	payload := new(dto.UpdateUserRequest)
	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	if err := c.Validate(payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	userId := c.Request().Context().Value("userId")

	result, err := h.service.Update(c.Request().Context(), uint(id), userId.(uint), payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status": result,
	})
}

func (h *handler) Delete(c echo.Context) error {
	payload := new(dto.ByIdRequest)
	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	if err := c.Validate(payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	userId := c.Request().Context().Value("userId")

	result, err := h.service.Delete(c.Request().Context(), userId.(uint), payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"deleted_user": result,
	})
}
