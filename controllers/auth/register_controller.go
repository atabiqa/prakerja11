package auth

import (
	"atabiqa/prakerja11/configs"
	"atabiqa/prakerja11/middleware"
	"atabiqa/prakerja11/models/base"
	usermodel "atabiqa/prakerja11/models/user"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func RegisterController(c echo.Context) error {
	var user usermodel.User
	c.Bind(&user)

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 13)
	user.Password = string(hashPassword)

	result := configs.DB.Create(&user)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, base.BaseResponse{
			Status:  false,
			Message: "failed add data to database",
			Data:    nil,
		})
	}

	var authRespon = usermodel.ResponseAuth{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
		Token: middleware.GenerateTokenJWT(user.Id, user.Name),
	}

	return c.JSON(http.StatusOK, base.BaseResponse{
		Status:  true,
		Message: "register succes",
		Data:    authRespon,
	})
}
