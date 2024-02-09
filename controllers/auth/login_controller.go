package auth

import (
	"atabiqa/prakerja11/models/base"
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoginController(c echo.Context) error {
	return c.JSON(http.StatusOK, base.BaseResponse{
		Status:  true,
		Message: "login succes",
		Data:    nil,
	})
}
