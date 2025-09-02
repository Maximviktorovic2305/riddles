package handlers

import (
	"net/http"
	"time"
	"riddles-server/services"

	"github.com/labstack/echo/v4"
)

type DailyRiddleHandler struct {
	dailyRiddleService services.DailyRiddleService
}

func NewDailyRiddleHandler(dailyRiddleService services.DailyRiddleService) *DailyRiddleHandler {
	return &DailyRiddleHandler{
		dailyRiddleService: dailyRiddleService,
	}
}

func (h *DailyRiddleHandler) GetTodayRiddle(c echo.Context) error {
	dailyRiddle, err := h.dailyRiddleService.GetTodayRiddle()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch today's riddle")
	}

	if dailyRiddle == nil {
		return echo.NewHTTPError(http.StatusNotFound, "No riddle found for today")
	}

	return c.JSON(http.StatusOK, dailyRiddle)
}

func (h *DailyRiddleHandler) GetRiddleByDate(c echo.Context) error {
	dateStr := c.Param("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid date format. Use YYYY-MM-DD")
	}

	dailyRiddle, err := h.dailyRiddleService.GetRiddleByDate(date)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch riddle for date")
	}

	if dailyRiddle == nil {
		return echo.NewHTTPError(http.StatusNotFound, "No riddle found for date")
	}

	return c.JSON(http.StatusOK, dailyRiddle)
}