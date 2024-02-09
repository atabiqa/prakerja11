package routes

import (
	"atabiqa/prakerja11/controllers/auth"
	"atabiqa/prakerja11/controllers/userr"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.POST("/login", auth.LoginController)
	e.POST("/register", auth.RegisterController)

	eAuth := e.Group("")
	eAuth.Use(echojwt.JWT([]byte("")))
	eAuth.GET("/users", userr.GetUserController)

}
