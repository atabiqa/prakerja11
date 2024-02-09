package userr

import (
	"atabiqa/prakerja11/configs"
	"atabiqa/prakerja11/models/base"
	"atabiqa/prakerja11/models/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUserController(e echo.Context) error {
	var users []user.User
	result := configs.DB.Find(&users)
	if result.Error != nil {
		return e.JSON(http.StatusInternalServerError, base.BaseResponse{
			Status:  false,
			Message: "failed get database",
			Data:    nil,
		})
	}
	return e.JSON(http.StatusOK, base.BaseResponse{
		Status:  true,
		Message: "succes",
		Data:    users,
	})
}
