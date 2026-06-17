package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RoleMiddleware(role string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			currentRole, ok := c.Get("role").(string)
			if !ok {
				return c.JSON(
					http.StatusForbidden,
					echo.Map{
						"message": "forbidden",
					},
				)
			}

			if currentRole != role {
				return c.JSON(
					http.StatusForbidden,
					echo.Map{
						"message": "forbidden",
					},
				)
			}

			return next(c)
		}
	}
}
