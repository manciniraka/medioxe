package helper

import "github.com/labstack/echo/v4"

func GetUserID(c echo.Context) int {
	return c.Get("user_id").(int)
}
