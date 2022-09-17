package controllers_test

import (
	"day3-middleware/config"
	"day3-middleware/helpers"
	"day3-middleware/routes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

func TestGetBooksController(t *testing.T) {
	config.InitDB()

	server := httptest.NewServer(routes.New())
	defer server.Close()

	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  server.URL,
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})

	t.Cleanup(func() {
		helpers.CleanBooksTable()
	})

	t.Run("should return 200", func(t *testing.T) {
		e.GET("/v1/books").Expect().Status(http.StatusOK)
	})
}

func TestCreateBookController(t *testing.T) {
	config.InitDB()

	server := httptest.NewServer(routes.New())
	defer server.Close()

	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  server.URL,
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})

	t.Cleanup(func() {
		helpers.CleanBooksTable()
		helpers.CleanUsersTable()
	})

	t.Run("should return 401 if not authenticated", func(t *testing.T) {
		e.POST("/v1/books").WithHeader("Authorization", "Bearer <bad token>").
			WithJSON(map[string]interface{}{
				"title":  "agmc",
				"author": "alta",
			}).Expect().Status(http.StatusUnauthorized)
	})

	t.Run("should return 400 if not meet payload requirements", func(t *testing.T) {
		helpers.AddUserWithEmailPassword("alta@gmail.com", "alta")
		token := helpers.GetJwtToken("alta@gmail.com", "alta")

		e.POST("/v1/books").WithHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
			WithJSON(map[string]interface{}{
				"title":  "agmc",
				"author": "alta",
			}).Expect().Status(http.StatusBadRequest)
	})

	t.Run("should return 200", func(t *testing.T) {
		helpers.AddUserWithEmailPassword("alta@gmail.com", "alta")
		token := helpers.GetJwtToken("alta@gmail.com", "alta")

		e.POST("/v1/books").WithHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
			WithJSON(map[string]interface{}{
				"title":  "agmc",
				"author": "alta",
				"year":   2022,
			}).Expect().Status(http.StatusCreated)
	})
}

func TestGetBookByIdController(t *testing.T) {
	config.InitDB()

	server := httptest.NewServer(routes.New())
	defer server.Close()

	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  server.URL,
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})

	t.Cleanup(func() {
		helpers.CleanBooksTable()
	})

	t.Run("should return 400 if book not found", func(t *testing.T) {
		e.GET("/v1/books/25").Expect().Status(http.StatusBadRequest)
	})

	t.Run("should return 200", func(t *testing.T) {
		helpers.AddBookWithId(25)

		e.GET("/v1/books/25").Expect().Status(http.StatusOK)
	})
}

func TestUpdateBookByIdController(t *testing.T) {
	config.InitDB()

	server := httptest.NewServer(routes.New())
	defer server.Close()

	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  server.URL,
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})

	t.Cleanup(func() {
		helpers.CleanBooksTable()
		helpers.CleanUsersTable()
	})

	t.Run("should return 401 if not authenticated", func(t *testing.T) {
		e.PUT("/v1/books/25").WithHeader("Authorization", "Bearer <bad token>").
			WithJSON(map[string]interface{}{
				"title":  "agmc",
				"author": "alta",
			}).Expect().Status(http.StatusUnauthorized)
	})

	t.Run("should return 400 if not meet payload requirements", func(t *testing.T) {
		helpers.AddUserWithEmailPassword("alta@gmail.com", "alta")
		token := helpers.GetJwtToken("alta@gmail.com", "alta")

		e.PUT("/v1/books/25").WithHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
			WithJSON(map[string]interface{}{
				"title":  "agmc",
				"author": "alta",
			}).Expect().Status(http.StatusBadRequest)
	})

	t.Run("should return 400 if book not found", func(t *testing.T) {
		helpers.AddUserWithEmailPassword("alta@gmail.com", "alta")
		token := helpers.GetJwtToken("alta@gmail.com", "alta")

		e.PUT("/v1/books/25").WithHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
			WithJSON(map[string]interface{}{
				"title":  "agmc",
				"author": "alta",
				"year":   2022,
			}).Expect().Status(http.StatusBadRequest)
	})

	t.Run("should return 200", func(t *testing.T) {
		helpers.AddBookWithId(25)
		helpers.AddUserWithEmailPassword("alta@gmail.com", "alta")
		token := helpers.GetJwtToken("alta@gmail.com", "alta")

		e.PUT("/v1/books/25").WithHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
			WithJSON(map[string]interface{}{
				"title":  "agmc",
				"author": "alta",
				"year":   2022,
			}).Expect().Status(http.StatusOK)
	})
}

func TestDeleteBookByIdHandler(t *testing.T) {
	config.InitDB()

	server := httptest.NewServer(routes.New())
	defer server.Close()

	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  server.URL,
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})

	t.Cleanup(func() {
		helpers.CleanBooksTable()
		helpers.CleanUsersTable()
	})

	t.Run("should return 401 if not authenticated", func(t *testing.T) {
		e.DELETE("/v1/books/25").WithHeader("Authorization", "Bearer <bad token>").
			Expect().Status(http.StatusUnauthorized)
	})

	t.Run("should return 200", func(t *testing.T) {
		helpers.AddBookWithId(25)

		helpers.AddUserWithEmailPassword("alta@gmail.com", "alta")
		token := helpers.GetJwtToken("alta@gmail.com", "alta")

		e.DELETE("/v1/books/25").WithHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
			Expect().Status(http.StatusOK)
	})
}
