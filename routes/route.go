package routes

import (
	"atabiqa/prakerja11/controllers/userr"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())

	e.GET("/users", userr.GetUserController)
}
