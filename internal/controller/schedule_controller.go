package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/manciniraka/medioxe/internal/helper"
	"github.com/manciniraka/medioxe/internal/service"
)

type ScheduleController struct {
	scheduleService service.ScheduleService
}

func NewScheduleController(
	scheduleService service.ScheduleService,
) *ScheduleController {
	return &ScheduleController{
		scheduleService: scheduleService,
	}
}

func (sc *ScheduleController) CreateSchedule(c echo.Context) error {
	var input service.CreateScheduleInput

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

	userID := helper.GetUserID(c)

	schedule, err := sc.scheduleService.CreateSchedule(userID, input)
	if err != nil {
		return helper.InternalServerError(
			c,
			err,
		)
	}

	return c.JSON(
		http.StatusCreated,
		echo.Map{
			"message": "schedule created successfully",
			"data":    schedule,
		},
	)
}

func (sc *ScheduleController) GetMySchedules(c echo.Context) error {
	userID := helper.GetUserID(c)

	schedules, err := sc.scheduleService.GetMySchedules(userID)
	if err != nil {
		return helper.InternalServerError(
			c,
			err,
		)
	}

	return c.JSON(
		http.StatusOK,
		echo.Map{
			"message": "schedule fetched successfully",
			"data":    schedules,
		},
	)
}

func (sc *ScheduleController) GetDoctorSchedules(c echo.Context) error {
	doctorID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			echo.Map{
				"message": "invalid doctor id",
			},
		)
	}

	schedule, err := sc.scheduleService.GetDoctorSchedules(doctorID)
	if err != nil {
		return helper.InternalServerError(
			c,
			err,
		)
	}

	return c.JSON(
		http.StatusOK,
		echo.Map{
			"message": "doctor fetched successfully",
			"data":    schedule,
		},
	)
}

func (sc *ScheduleController) UpdateSchedule(c echo.Context) error {
	scheduleID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			echo.Map{
				"message": "invalid schedule id",
			},
		)
	}

	var input service.UpdateScheduleInput

	if err := c.Bind(&input); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			echo.Map{
				"message": "invalid request",
			},
		)
	}

	userID := helper.GetUserID(c)

	schedule, err := sc.scheduleService.UpdateSchedule(
		userID,
		scheduleID,
		input,
	)
	if err != nil {
		return helper.InternalServerError(
			c,
			err,
		)
	}

	return c.JSON(
		http.StatusOK,
		echo.Map{
			"message": "schedule updated successfully",
			"data":    schedule,
		},
	)
}

func (sc *ScheduleController) DeleteSchedule(c echo.Context) error {
	scheduleID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			echo.Map{
				"message": "invalid schedule id",
			},
		)
	}

	userID := helper.GetUserID(c)

	err = sc.scheduleService.DeleteSchedule(
		userID,
		scheduleID,
	)
	if err != nil {
		return helper.InternalServerError(
			c,
			err,
		)
	}

	return c.JSON(
		http.StatusOK,
		echo.Map{
			"message": "schedule deleted successfully",
		},
	)
}
