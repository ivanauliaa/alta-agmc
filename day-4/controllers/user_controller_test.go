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

func TestGetUsersController(t *testing.T) {
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
		helpers.CleanUsersTable()
	})

	t.Run("should return 401 if not authenticated", func(t *testing.T) {
		e.GET("/v1/users").WithHeader("Authorization", "Bearer <bad token>").
			Expect().Status(http.StatusUnauthorized)
	})

	t.Run("should return 200 if authenticated", func(t *testing.T) {
		helpers.AddUserWithEmailPassword("alta@gmail.com", "alta")
		token := helpers.GetJwtToken("alta@gmail.com", "alta")

		e.GET("/v1/users").WithHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
			Expect().Status(http.StatusOK)
	})
}

func TestCreateUsersController(t *testing.T) {
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
		helpers.CleanUsersTable()
	})

	t.Run("should return error if not meet payload requirements", func(t *testing.T) {
		e.POST("/v1/users").WithJSON(map[string]interface{}{
			"name":  "agmc",
			"email": "agmc@gmail.com",
		}).Expect().Status(http.StatusBadRequest)
	})

	t.Run("should return 201 if meet payload requirements", func(t *testing.T) {
		e.POST("/v1/users").WithJSON(map[string]interface{}{
			"name":     "agmc",
			"email":    "agmc@gmail.com",
			"password": "agmc",
		}).Expect().Status(http.StatusCreated)
	})
}

func TestGetUserByIdController(t *testing.T) {
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
		helpers.CleanUsersTable()
	})

	t.Run("should return 401 if not authenticated", func(t *testing.T) {
		e.GET("/v1/users/20").WithHeader("Authorization", "Bearer <bad token>").
			Expect().Status(http.StatusUnauthorized)
	})

	t.Run("should return 400 if user not found", func(t *testing.T) {
		helpers.AddUserWithEmailPassword("alta@gmail.com", "alta")
		token := helpers.GetJwtToken("alta@gmail.com", "alta")

		e.GET("/v1/users/999").WithHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
			Expect().Status(http.StatusBadRequest)
	})

	t.Run("should return 200 if user found", func(t *testing.T) {
		helpers.AddUserWithId(999)

		helpers.AddUserWithEmailPassword("alta@gmail.com", "alta")
		token := helpers.GetJwtToken("alta@gmail.com", "alta")

		e.GET("/v1/users/999").WithHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
			Expect().Status(http.StatusOK)
	})
}

func TestUpdateUserByIdController(t *testing.T) {
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
		helpers.CleanUsersTable()
	})

	t.Run("should return 401 if not authenticated", func(t *testing.T) {
		e.PUT("/v1/users/20").WithHeader("Authorization", "Bearer <bad token>").
			Expect().Status(http.StatusUnauthorized)
	})

	t.Run("should return 403 if access restricted resource", func(t *testing.T) {
		helpers.AddUserWithEmailPassword("alta@gmail.com", "alta")
		token := helpers.GetJwtToken("alta@gmail.com", "alta")

		e.PUT("/v1/users/999").WithHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
			Expect().Status(http.StatusForbidden)
	})

	t.Run("should return 400 if not meet payload requirements", func(t *testing.T) {
		helpers.AddUserWithIdEmailPassword(25, "alta@gmail.com", "alta")
		token := helpers.GetJwtToken("alta@gmail.com", "alta")

		e.PUT("/v1/users/25").WithHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
			WithJSON(map[string]interface{}{
				"name":  "agmc",
				"email": "agmc@gmail.com",
			}).Expect().Status(http.StatusBadRequest)
	})

	t.Run("should return 200 if meet payload requirements", func(t *testing.T) {
		helpers.AddUserWithIdEmailPassword(25, "alta@gmail.com", "alta")
		token := helpers.GetJwtToken("alta@gmail.com", "alta")

		e.PUT("/v1/users/25").WithHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
			WithJSON(map[string]interface{}{
				"name":     "agmc",
				"email":    "agmc@gmail.com",
				"password": "agmc",
			}).Expect().Status(http.StatusOK)
	})
}

func TestDeleteUserByIdController(t *testing.T) {
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
		helpers.CleanUsersTable()
	})

	t.Run("should return 401 if not authenticated", func(t *testing.T) {
		e.DELETE("/v1/users/20").WithHeader("Authorization", "Bearer <bad token>").
			Expect().Status(http.StatusUnauthorized)
	})

	t.Run("should return 403 if access restricted resource", func(t *testing.T) {
		helpers.AddUserWithEmailPassword("alta@gmail.com", "alta")
		token := helpers.GetJwtToken("alta@gmail.com", "alta")

		e.DELETE("/v1/users/999").WithHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
			Expect().Status(http.StatusForbidden)
	})

	t.Run("should return 200", func(t *testing.T) {
		helpers.AddUserWithIdEmailPassword(25, "alta@gmail.com", "alta")
		token := helpers.GetJwtToken("alta@gmail.com", "alta")

		e.DELETE("/v1/users/25").WithHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
			Expect().Status(http.StatusOK)
	})
}

func TestLoginController(t *testing.T) {
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
		helpers.CleanUsersTable()
	})

	t.Run("should return 400 if user not found", func(t *testing.T) {
		e.POST("/v1/login").WithJSON(map[string]interface{}{
			"email":    "alta@gmail.com",
			"password": "alta",
		}).Expect().Status(http.StatusBadRequest)
	})

	t.Run("should return 200 if user found", func(t *testing.T) {
		helpers.AddUserWithEmailPassword("alta@gmail.com", "alta")

		e.POST("/v1/login").WithJSON(map[string]interface{}{
			"email":    "alta@gmail.com",
			"password": "alta",
		}).Expect().Status(http.StatusOK)
	})
}
