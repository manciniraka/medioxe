package router

import (
	"github.com/labstack/echo/v4"
	"github.com/manciniraka/medioxe/internal/controller"
	"github.com/manciniraka/medioxe/internal/middleware"
	"github.com/manciniraka/medioxe/internal/repository"
	"github.com/manciniraka/medioxe/internal/service"
	"gorm.io/gorm"
)

func InitRouter(
	e *echo.Echo,
	db *gorm.DB,
) {
	userRepo := repository.NewUserRepository(db)

	authService := service.NewAuthService(userRepo)
	authController := controller.NewAuthController(authService)

	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	e.POST("/register", authController.Register)
	e.POST("/login", authController.Login)

	auth := e.Group("")

	auth.Use(middleware.AuthMiddleware)

	auth.GET("/profile", userController.GetProfile)

}
