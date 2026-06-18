package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/manciniraka/medioxe/internal/helper"
	"github.com/manciniraka/medioxe/internal/service"
)

type AIController struct {
	aiService service.AIService
}

func NewAIController(aiService service.AIService) *AIController {
	return &AIController{
		aiService: aiService,
	}
}

func (ac *AIController) AnalyzeSymptoms(c echo.Context) error {
	var input service.SymptomAnalysisInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			echo.Map{
				"message": "invalid request body",
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

	result, err := ac.aiService.AnalyzeSymptoms(
		userID,
		input,
	)

	if err != nil {
		if err.Error() == "specialty not found" {
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
			"message": "symptom analysis completed",
			"data":    result,
		},
	)

}
