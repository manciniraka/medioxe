package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/manciniraka/medioxe/internal/helper"
	"github.com/manciniraka/medioxe/internal/service"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) GetProfile(c echo.Context) error {
	userID := helper.GetUserID(c)

	user, err := uc.userService.GetProfile(userID)
	if err != nil {
		return c.JSON(
			http.StatusNotFound,
			echo.Map{
				"message": "user not found",
			},
		)
	}

	return c.JSON(
		http.StatusOK,
		echo.Map{
			"message": "get profile data success",
			"data":    user,
		},
	)
}
