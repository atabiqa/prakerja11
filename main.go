package main

import (
	"atabiqa/prakerja11/configs"
	"atabiqa/prakerja11/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	configs.InitDatabase()
	e := echo.New()
	routes.InitRoute(e)
	e.Start(":7575")
}
