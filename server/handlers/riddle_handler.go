package handlers

import (
	"net/http"
	"strconv"
	"riddles-server/services"

	"github.com/labstack/echo/v4"
)

type RiddleHandler struct {
	riddleService services.RiddleService
}

func NewRiddleHandler(riddleService services.RiddleService) *RiddleHandler {
	return &RiddleHandler{
		riddleService: riddleService,
	}
}

type CheckAnswerRequest struct {
	Answer string `json:"answer" validate:"required"`
}

type CheckAnswerResponse struct {
	Correct bool   `json:"correct"`
	Message string `json:"message"`
}

func (h *RiddleHandler) GetAllRiddles(c echo.Context) error {
	riddles, err := h.riddleService.GetAllRiddles()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch riddles")
	}

	// TODO: Extract user ID from JWT token for personalized data
	// For now, we'll use a placeholder
	userID := uint(1)

	riddlesWithProgress, err := h.riddleService.GetRiddlesWithUserProgress(riddles, userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch riddles with progress")
	}

	return c.JSON(http.StatusOK, riddlesWithProgress)
}

func (h *RiddleHandler) GetRiddleByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid riddle ID")
	}

	// TODO: Extract user ID from JWT token for personalized data
	// For now, we'll use a placeholder
	userID := uint(1)

	riddleWithProgress, err := h.riddleService.GetRiddleWithUserProgress(uint(id), userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Riddle not found")
	}

	return c.JSON(http.StatusOK, riddleWithProgress)
}

func (h *RiddleHandler) CheckAnswer(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid riddle ID")
	}

	var req CheckAnswerRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	correct, err := h.riddleService.CheckAnswer(uint(id), req.Answer)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Riddle not found")
	}

	response := CheckAnswerResponse{
		Correct: correct,
	}

	if correct {
		response.Message = "Правильный ответ! Поздравляем!"
	} else {
		response.Message = "Неправильный ответ. Попробуйте еще раз!"
	}

	return c.JSON(http.StatusOK, response)
}