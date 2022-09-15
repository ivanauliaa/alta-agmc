package controllers

import (
	"day3-middleware/lib/database"
	"day3-middleware/middlewares"
	"day3-middleware/models"
	"net/http"
	"strconv"

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
	bindUser := models.User{}
	if err := c.Bind(&bindUser); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(bindUser); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := database.CreateUser(bindUser)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create user",
		"user":    user,
	})
}

func GetUserByIdController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := database.GetUserById(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user",
		"user":    user,
	})
}

func UpdateUserByIdController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	owner := middlewares.ExtractTokenUserId(&c)
	if err := database.VerifyUserOwner(uint(id), owner); err != nil {
		return echo.NewHTTPError(http.StatusForbidden, err.Error())
	}

	var userUpdate models.User

	if err := c.Bind(&userUpdate); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(userUpdate); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := database.UpdateUserById(uint(id), userUpdate)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
		"user":    user,
	})
}

func DeleteUserByIdController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	owner := middlewares.ExtractTokenUserId(&c)
	if err := database.VerifyUserOwner(uint(id), owner); err != nil {
		return echo.NewHTTPError(http.StatusForbidden, err.Error())
	}

	if err := database.DeleteUserById(uint(id)); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
	})
}

func LoginController(c echo.Context) error {
	bindUser := models.UserLogin{}
	if err := c.Bind(&bindUser); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(bindUser); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := database.LoginUser(bindUser)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	token, err := middlewares.CreateToken(user.(models.User).ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":      "success login",
		"access_token": token,
	})
}
