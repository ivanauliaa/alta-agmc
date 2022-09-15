package main

import (
	"day3-middleware/config"
	"day3-middleware/routes"
)

func main() {
	config.InitDB()

	e := routes.New()

	e.Logger.Fatal(e.Start(":5000"))
}
