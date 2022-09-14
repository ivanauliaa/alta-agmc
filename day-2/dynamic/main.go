package main

import (
	"day2-dynamic/config"
	"day2-dynamic/routes"
)

func main() {
	config.InitDB()

	e := routes.New()

	e.Logger.Fatal(e.Start(":5000"))
}
