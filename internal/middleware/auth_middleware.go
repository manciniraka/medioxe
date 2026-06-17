package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	commonJWT "github.com/manciniraka/go-common/jwt"
)

func AuthMiddleware(
	next echo.HandlerFunc,
) echo.HandlerFunc {

	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")

		if authHeader == "" {
			return c.JSON(
				http.StatusUnauthorized,
				echo.Map{
					"message": "missing token",
				},
			)
		}

		splitToken := strings.Split(authHeader, " ")
		if len(splitToken) != 2 {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": "invalid token",
			})
		}

		tokenString := splitToken[1]

		// Parsing token using my personal module packages
		claims, err := commonJWT.ParseToken(
			tokenString,
			os.Getenv("JWT_SECRET"),
		)
		if err != nil {
			return c.JSON(
				http.StatusUnauthorized,
				echo.Map{
					"message": "invalid token",
				},
			)
		}

		userID := int(claims["user_id"].(float64))
		c.Set("user_id", userID)
		c.Set("role", claims["role"])

		return next(c)
	}
}
