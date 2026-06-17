package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/manciniraka/medioxe/internal/service"
)

type AuthController struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

// Creates a new user account
func (ac *AuthController) Register(c echo.Context) error {
	var input service.RegisterInput

	if err := c.Bind(&input); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			echo.Map{
				"message": "invalid request body",
			},
		)
	}
	if err := c.Validate(&input); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			echo.Map{
				"message": err.Error(),
			},
		)
	}

	user, err := ac.authService.Register(input)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			echo.Map{
				"message": err.Error(),
			},
		)
	}

	return c.JSON(
		http.StatusCreated,
		echo.Map{
			"message": "user registered successfully",
			"data":    user,
		},
	)
}

// Authenticates user and returns JWT Token
func (ac *AuthController) Login(c echo.Context) error {
	var input service.LoginInput

	if err := c.Bind(&input); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			echo.Map{
				"message": "invalid request",
			},
		)
	}
	if err := c.Validate(&input); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			echo.Map{
				"message": err.Error(),
			},
		)
	}

	token, err := ac.authService.Login(input)
	if err != nil {
		return c.JSON(
			http.StatusUnauthorized,
			echo.Map{
				"message": err.Error(),
			},
		)
	}

	return c.JSON(
		http.StatusOK,
		echo.Map{
			"message":      "login successfully",
			"access_token": token,
		},
	)
}
