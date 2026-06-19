package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/manciniraka/medioxe/internal/helper"
	"github.com/manciniraka/medioxe/internal/service"
)

type AppointmentController struct {
	appointmentService service.AppointmentService
}

func NewAppointmentController(
	appointmentService service.AppointmentService,
) *AppointmentController {
	return &AppointmentController{
		appointmentService: appointmentService,
	}
}

func (ac *AppointmentController) CreateAppointment(c echo.Context) error {
	var input service.CreateAppointmentInput

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

	patientID := helper.GetUserID(c)

	appointment, err := ac.appointmentService.CreateAppointment(patientID, input)
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
			"message": "appointment created successfully",
			"data":    appointment,
		},
	)
}

func (ac *AppointmentController) GetMyAppointments(c echo.Context) error {
	patientID := helper.GetUserID(c)

	appointment, err := ac.appointmentService.GetMyAppointments(patientID)
	if err != nil {
		return helper.InternalServerError(
			c,
			err,
		)
	}

	return c.JSON(
		http.StatusOK,
		echo.Map{
			"message": "appointment fetched successfully",
			"data":    appointment,
		},
	)
}

func (ac *AppointmentController) GetDoctorAppointments(c echo.Context) error {
	userID := helper.GetUserID(c)

	appointment, err := ac.appointmentService.GetDoctorAppointments(userID)
	if err != nil {
		return helper.InternalServerError(
			c,
			err,
		)
	}

	return c.JSON(
		http.StatusOK,
		echo.Map{
			"message": "appointment fetched successfully",
			"data":    appointment,
		},
	)
}

func (ac *AppointmentController) ConfirmAppointment(c echo.Context) error {
	appointmentID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			echo.Map{
				"message": "invalid appointment id",
			},
		)
	}

	userID := helper.GetUserID(c)

	err = ac.appointmentService.ConfirmAppointment(userID, appointmentID)
	if err != nil {
		if err.Error() == "appointment not found" ||
			err.Error() == "appointment already processed" ||
			err.Error() == "you are not allowed to access this appointment" {
			return c.JSON(
				http.StatusBadRequest,
				echo.Map{
					"message": err.Error(),
				},
			)
		}

		return helper.InternalServerError(
			c,
			err,
		)
	}

	return c.JSON(
		http.StatusOK,
		echo.Map{
			"message": "appointment confirmed successfully",
		},
	)
}

func (ac *AppointmentController) CompleteAppointment(c echo.Context) error {
	appointmentID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			echo.Map{
				"message": "invalid appointment id",
			},
		)
	}

	userID := helper.GetUserID(c)

	err = ac.appointmentService.CompleteAppointment(userID, appointmentID)
	if err != nil {
		if err.Error() == "appointment not found" ||
			err.Error() == "appointment already processed" ||
			err.Error() == "appointment must be confirmed first" ||
			err.Error() == "you are not allowed to access this appointment" {
			return c.JSON(
				http.StatusBadRequest,
				echo.Map{
					"message": err.Error(),
				},
			)
		}

		return helper.InternalServerError(
			c,
			err,
		)
	}

	return c.JSON(
		http.StatusOK,
		echo.Map{
			"message": "appointment Completed successfully",
		},
	)
}

func (ac *AppointmentController) CancelAppointment(c echo.Context) error {
	appointmentID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			echo.Map{
				"message": "invalid appointment id",
			},
		)
	}

	userID := helper.GetUserID(c)

	err = ac.appointmentService.CancelAppointment(userID, appointmentID)
	if err != nil {
		if err.Error() == "appointment not found" ||
			err.Error() == "appointment already processed" ||
			err.Error() == "completed appointment cannot be cancelled" ||
			err.Error() == "you are not allowed to access this appointment" {
			return c.JSON(
				http.StatusBadRequest,
				echo.Map{
					"message": err.Error(),
				},
			)
		}

		return helper.InternalServerError(
			c,
			err,
		)
	}

	return c.JSON(
		http.StatusOK,
		echo.Map{
			"message": "appointment cancelled successfully",
		},
	)
}

func (ac *AppointmentController) GetAppointmentHistory(c echo.Context) error {
	appointmentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			echo.Map{
				"message": "invalid appointment id",
			},
		)
	}

	history, err := ac.appointmentService.GetAppointmentHistory(appointmentID)
	if err != nil {
		return helper.InternalServerError(
			c,
			err,
		)
	}

	return c.JSON(
		http.StatusOK,
		echo.Map{
			"message": "appointment history fetched successfully",
			"data":    history,
		},
	)
}

func (ac *AppointmentController) GetAllAppointmentHistories(c echo.Context) error {
	histories, err := ac.appointmentService.GetAll()

	if err != nil {
		return helper.InternalServerError(
			c,
			err,
		)
	}

	return c.JSON(
		http.StatusOK,
		echo.Map{
			"message": "appointment histories fetched successfully",
			"data":    histories,
		},
	)
}
