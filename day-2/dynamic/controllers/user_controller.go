package controllers

import (
	"day2-dynamic/lib/database"
	"day2-dynamic/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsersController(c echo.Context) error {
	users, err := database.GetUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

func CreateUserController(c echo.Context) error {
	userBind := models.User{}
	c.Bind(&userBind)

	user, err := database.CreateUser(userBind)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create user",
		"user":    user,
	})
}

func GetUserByIdController(c echo.Context) error {
	id := c.Param("id")

	user, err := database.GetUserById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user",
		"user":    user,
	})
}

func UpdateUserByIdController(c echo.Context) error {
	id := c.Param("id")
	var userUpdate models.User

	if err := c.Bind(&userUpdate); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := database.UpdateUserById(id, userUpdate)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
		"user":    user,
	})
}

func DeleteUserByIdController(c echo.Context) error {
	id := c.Param("id")

	if err := database.DeleteUserById(id); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
	})
}
