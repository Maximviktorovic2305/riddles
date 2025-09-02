package handlers

import (
	"net/http"
	"strconv"
	"riddles-server/services"

	"github.com/labstack/echo/v4"
)

type RatingHandler struct {
	ratingService services.RatingService
}

type RateRiddleRequest struct {
	Rating int `json:"rating" validate:"required"`
}

func NewRatingHandler(ratingService services.RatingService) *RatingHandler {
	return &RatingHandler{
		ratingService: ratingService,
	}
}

func (h *RatingHandler) RateRiddle(c echo.Context) error {
	riddleID, err := strconv.Atoi(c.Param("riddle_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid riddle ID")
	}

	var req RateRiddleRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	// TODO: Extract user ID from JWT token
	// For now, we'll use a placeholder
	userID := uint(1)

	err = h.ratingService.RateRiddle(userID, uint(riddleID), req.Rating)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to rate riddle")
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *RatingHandler) RemoveRating(c echo.Context) error {
	riddleID, err := strconv.Atoi(c.Param("riddle_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid riddle ID")
	}

	// TODO: Extract user ID from JWT token
	// For now, we'll use a placeholder
	userID := uint(1)

	err = h.ratingService.RemoveRating(userID, uint(riddleID))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to remove rating")
	}

	return c.NoContent(http.StatusNoContent)
}