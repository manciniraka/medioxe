package helper

import (
	"github.com/labstack/echo/v4"
	commonLogger "github.com/manciniraka/go-common/logger"
)

// Logger using my personal module packages
var appLogger = commonLogger.NewJSON()

func InternalServerError(c echo.Context, err error) error {
	appLogger.Error(err.Error())

	return c.JSON(
		500,
		echo.Map{
			"message": "internal server error",
		},
	)
}
