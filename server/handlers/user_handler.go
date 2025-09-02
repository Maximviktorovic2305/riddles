package handlers

import (
	"net/http"
	"strconv"
	"riddles-server/services"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

type UserStatsResponse struct {
	TotalRiddles   int `json:"total_riddles"`
	SolvedRiddles  int `json:"solved_riddles"`
	SuccessRate    int `json:"success_rate"` // percentage
}

func (h *UserHandler) GetProfile(c echo.Context) error {
	// TODO: Extract user ID from JWT token
	// For now, we'll use a placeholder
	userID := uint(1) // This should come from the authenticated user

	user, err := h.userService.GetProfile(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetUserStats(c echo.Context) error {
	// TODO: Extract user ID from JWT token
	// For now, we'll use a placeholder
	userID := uint(1) // This should come from the authenticated user

	total, solved, err := h.userService.GetUserStats(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get user stats")
	}

	successRate := 0
	if total > 0 {
		successRate = (solved * 100) / total
	}

	return c.JSON(http.StatusOK, UserStatsResponse{
		TotalRiddles:  total,
		SolvedRiddles: solved,
		SuccessRate:   successRate,
	})
}