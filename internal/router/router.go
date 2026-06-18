package router

import (
	"github.com/labstack/echo/v4"
	geminiai "github.com/manciniraka/medioxe/external/geminiAI"
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
	specialtyRepo := repository.NewSpecialtyRepository(db)
	doctorRepo := repository.NewDoctorRepository(db)
	scheduleRepo := repository.NewScheduleRepository(db)
	symptomAnalysisRepo := repository.NewSymptomAnalysisRepository(db)

	authService := service.NewAuthService(userRepo)
	authController := controller.NewAuthController(authService)

	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	doctorService := service.NewDoctorService(doctorRepo, userRepo)
	doctorController := controller.NewDoctorController(doctorService)

	geminiClient := geminiai.NewGeminiClient()

	aiService := service.NewAIService(
		geminiClient,
		specialtyRepo,
		doctorRepo,
		symptomAnalysisRepo,
		scheduleRepo,
	)

	aiController := controller.NewAIController(aiService)

	e.POST("/register", authController.Register)
	e.POST("/login", authController.Login)

	auth := e.Group("")

	auth.Use(middleware.AuthMiddleware)

	auth.GET("/profile", userController.GetProfile)
	auth.GET("/doctors", doctorController.GetDoctors)

	auth.GET("/doctors/:id", doctorController.GetDoctorByID)

	patient := e.Group("")

	patient.Use(
		middleware.AuthMiddleware,
		middleware.RoleMiddleware("patient"),
	)

	patient.POST("/geminiai/symptom-analysis", aiController.AnalyzeSymptoms)

	admin := e.Group("/admin")

	admin.Use(
		middleware.AuthMiddleware,
		middleware.RoleMiddleware("admin"),
	)

	admin.POST("/doctors", doctorController.CreateDoctor)

	admin.PUT("/doctors/:id", doctorController.UpdateDoctor)

	admin.PATCH("/doctors/:id/activate", doctorController.ActivateDoctor)

	admin.PATCH("/doctors/:id/deactivate", doctorController.DeactivateDoctor)
}
