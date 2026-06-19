package router

import (
	"os"

	"github.com/labstack/echo/v4"
	geminiai "github.com/manciniraka/medioxe/external/geminiAI"
	"github.com/manciniraka/medioxe/external/mailjet"
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
	appointmentRepo := repository.NewAppointmentRepository(db)
	appointmentHistoryRepo := repository.NewAppointmentHistoryRepository(db)
	hospitalRepo := repository.NewHospitalRepository(db)

	mailjetRepo := mailjet.NewMailjetRepository(
		mailjet.MailjetConfig{
			MailjetBaseURL:           os.Getenv("MAILJET_BASE_URL"),
			MailjetBasicAuthUsername: os.Getenv("MAILJET_API_KEY"),
			MailjetBasicAuthPassword: os.Getenv("MAILJET_SECRET_KEY"),
			MailjetSenderEmail:       os.Getenv("MAILJET_SENDER_EMAIL"),
			MailjetSenderName:        os.Getenv("MAILJET_SENDER_NAME"),
		},
	)

	authService := service.NewAuthService(userRepo)
	authController := controller.NewAuthController(authService)

	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	doctorService := service.NewDoctorService(
		doctorRepo,
		userRepo,
		specialtyRepo,
		hospitalRepo,
	)
	doctorController := controller.NewDoctorController(doctorService)

	scheduleService := service.NewScheduleService(scheduleRepo, doctorRepo)
	scheduleController := controller.NewScheduleController(scheduleService)

	appointmentService := service.NewAppointmentService(
		appointmentRepo,
		appointmentHistoryRepo,
		scheduleRepo,
		doctorRepo,
		userRepo,
		mailjetRepo,
	)
	appointmentController := controller.NewAppointmentController(appointmentService)

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

	auth.GET("/doctors/:id/schedules", scheduleController.GetDoctorSchedules)

	patient := e.Group("")

	patient.Use(
		middleware.AuthMiddleware,
		middleware.RoleMiddleware("patient"),
	)

	patient.POST("/geminiai/symptom-analysis", aiController.AnalyzeSymptoms)

	patient.POST("/appointments", appointmentController.CreateAppointment)

	patient.GET("/appointments", appointmentController.GetMyAppointments)
	patient.GET("/appointments/:id/history", appointmentController.GetAppointmentHistory)

	admin := e.Group("/admin")

	admin.Use(
		middleware.AuthMiddleware,
		middleware.RoleMiddleware("admin"),
	)

	admin.POST("/doctors", doctorController.CreateDoctor)

	admin.PUT("/doctors/:id", doctorController.UpdateDoctor)
	admin.PATCH("/doctors/:id/activate", doctorController.ActivateDoctor)
	admin.PATCH("/doctors/:id/deactivate", doctorController.DeactivateDoctor)

	admin.GET("/appointment-histories", appointmentController.GetAllAppointmentHistories)
	admin.GET("/appointments/:id/history", appointmentController.GetAppointmentHistory)

	doctor := e.Group("")
	doctor.Use(
		middleware.AuthMiddleware,
		middleware.RoleMiddleware("doctor"),
	)

	doctor.GET("/doctor/profile", doctorController.GetMyProfile)

	doctor.PUT("/doctor/profile", doctorController.UpdateMyProfile)

	doctor.POST("/doctor/schedules", scheduleController.CreateSchedule)

	doctor.GET("/doctor/schedules", scheduleController.GetMySchedules)

	doctor.PUT("/doctor/schedules/:id", scheduleController.UpdateSchedule)

	doctor.DELETE("/doctor/schedules/:id", scheduleController.DeleteSchedule)

	doctor.GET("/doctor/appointments", appointmentController.GetDoctorAppointments)

	doctor.PATCH("/doctor/appointments/:id/confirm", appointmentController.ConfirmAppointment)
	doctor.PATCH("/doctor/appointments/:id/complete", appointmentController.CompleteAppointment)
	doctor.PATCH("/doctor/appointments/:id/cancel", appointmentController.CancelAppointment)

}
