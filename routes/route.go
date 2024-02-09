package routes

import (
	"atabiqa/prakerja11/controllers/userr"

	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) {

	e.GET("/users", userr.GetUserController)
}
