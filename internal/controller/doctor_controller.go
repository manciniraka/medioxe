package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/manciniraka/medioxe/internal/helper"
	"github.com/manciniraka/medioxe/internal/service"
)

type DoctorController struct {
	doctorService service.DoctorService
}

func NewDoctorController(
	doctorService service.DoctorService,
) *DoctorController {
	return &DoctorController{
		doctorService: doctorService,
	}
}

func (dc *DoctorController) CreateDoctor(c echo.Context) error {
	var input service.CreateDoctorInput

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

	doctor, err := dc.doctorService.CreateDoctor(input)
	if err != nil {

		if err.Error() == "email already registered" {
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
		http.StatusCreated,
		echo.Map{
			"message": "doctor created successfully",
			"data":    doctor,
		},
	)
}

func (dc *DoctorController) UpdateDoctor(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			echo.Map{
				"message": "invalid doctor id",
			},
		)
	}

	var input service.UpdateDoctorInput

	if err := c.Bind(&input); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			echo.Map{
				"message": "invalid request",
			},
		)
	}

	doctor, err := dc.doctorService.UpdateDoctor(id, input)
	if err != nil {
		if err.Error() == "doctor not found" {
			return c.JSON(
				http.StatusNotFound,
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
			"message": "doctor updated successfully",
			"data":    doctor,
		},
	)
}

func (dc *DoctorController) ActivateDoctor(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			echo.Map{
				"message": "invalid doctor id",
			},
		)
	}

	err = dc.doctorService.ActivateDoctor(id)

	if err != nil {
		if err.Error() == "doctor not found" {
			return c.JSON(
				http.StatusNotFound,
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
			"message": "doctor activated successfully",
		},
	)
}

func (dc *DoctorController) DeactivateDoctor(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			echo.Map{
				"message": "invalid doctor id",
			},
		)
	}

	err = dc.doctorService.DeactivateDoctor(id)

	if err != nil {
		if err.Error() == "doctor not found" {
			return c.JSON(
				http.StatusNotFound,
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
			"message": "doctor deactivated successfully",
		},
	)
}

func (dc *DoctorController) GetDoctors(c echo.Context) error {
	specialtyID, _ := strconv.Atoi(c.QueryParam("specialty_id"))

	hospitalID, _ := strconv.Atoi(c.QueryParam("hospital_id"))

	doctors, err := dc.doctorService.GetDoctors(specialtyID, hospitalID)

	if err != nil {
		return helper.InternalServerError(
			c,
			err,
		)
	}

	return c.JSON(
		http.StatusOK,
		echo.Map{
			"message": "doctors fetched successfully",
			"data":    doctors,
		},
	)
}

func (dc *DoctorController) GetDoctorByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			echo.Map{
				"message": "invalid doctor id",
			},
		)
	}

	doctor, err := dc.doctorService.GetDoctorByID(id)

	if err != nil {
		return c.JSON(
			http.StatusNotFound,
			echo.Map{
				"message": "doctor not found",
			},
		)
	}

	return c.JSON(
		http.StatusOK,
		echo.Map{
			"message": "doctor fetched successfully",
			"data":    doctor,
		},
	)
}

func (dc *DoctorController) GetMyProfile(c echo.Context) error {
	userID := helper.GetUserID(c)

	doctor, err := dc.doctorService.GetMyProfile(userID)
	if err != nil {
		return helper.InternalServerError(c, err)
	}

	return c.JSON(
		http.StatusOK,
		echo.Map{
			"message": "doctor profile fetched succesfully",
			"data":    doctor,
		},
	)
}

func (dc *DoctorController) UpdateMyProfile(c echo.Context) error {
	var input service.UpdateMyProfileInput

	if err := c.Bind(&input); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			echo.Map{
				"message": "invalid request body",
			},
		)
	}

	userID := helper.GetUserID(c)

	doctor, err := dc.doctorService.UpdateMyProfile(userID, input)
	if err != nil {
		return helper.InternalServerError(c, err)
	}

	return c.JSON(
		http.StatusOK,
		echo.Map{
			"message": "doctor profile fetched succesfully",
			"data":    doctor,
		},
	)
}
