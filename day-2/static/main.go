package main

import (
	"day2-static/routes"
)

func main() {
	e := routes.New()

	e.Logger.Fatal(e.Start(":5000"))
}
