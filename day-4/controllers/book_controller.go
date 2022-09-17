package controllers

import (
	"day3-middleware/lib/database"
	"day3-middleware/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetBooksController(c echo.Context) error {
	books, err := database.GetBooks()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"books":   books,
	})
}

func CreateBookController(c echo.Context) error {
	bookBind := models.Book{}
	if err := c.Bind(&bookBind); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(bookBind); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	book, err := database.CreateBook(bookBind)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success create book",
		"book":    book,
	})
}

func GetBookByIdController(c echo.Context) error {
	id := c.Param("id")

	book, err := database.GetBookById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get book",
		"book":    book,
	})
}

func UpdateBookByIdController(c echo.Context) error {
	id := c.Param("id")
	var bookUpdate models.Book

	if err := c.Bind(&bookUpdate); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(bookUpdate); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	book, err := database.UpdateBookById(id, bookUpdate)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update book",
		"book":    book,
	})
}

func DeleteBookByIdController(c echo.Context) error {
	id := c.Param("id")

	if err := database.DeleteBookById(id); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete book",
	})
}
