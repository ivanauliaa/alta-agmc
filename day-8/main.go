package main

import (
	"os"

	"day6-hexagonal/database"
	"day6-hexagonal/database/migration"
	"day6-hexagonal/internal/factory"
	"day6-hexagonal/internal/http"
	"day6-hexagonal/internal/middleware"
	"day6-hexagonal/mongo"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	if os.Getenv("APP_ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			panic(err)
		}
	}
}

func main() {
	database.CreateConnection()
	mongo.CreateConnection()

	migration.Migrate()

	f := factory.NewFactory()
	e := echo.New()
	middleware.Init(e)
	http.NewHttp(e, f)

	e.Logger.Fatal(e.Start(":" + os.Getenv("APP_PORT")))
}
