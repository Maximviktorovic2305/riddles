package handlers

import (
	"net/http"
	"strconv"
	"riddles-server/services"

	"github.com/labstack/echo/v4"
)

type FavoriteHandler struct {
	favoriteService services.FavoriteService
}

func NewFavoriteHandler(favoriteService services.FavoriteService) *FavoriteHandler {
	return &FavoriteHandler{
		favoriteService: favoriteService,
	}
}

func (h *FavoriteHandler) AddFavorite(c echo.Context) error {
	riddleID, err := strconv.Atoi(c.Param("riddle_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid riddle ID")
	}

	// TODO: Extract user ID from JWT token
	// For now, we'll use a placeholder
	userID := uint(1)

	err = h.favoriteService.AddFavorite(userID, uint(riddleID))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to add favorite")
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *FavoriteHandler) RemoveFavorite(c echo.Context) error {
	riddleID, err := strconv.Atoi(c.Param("riddle_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid riddle ID")
	}

	// TODO: Extract user ID from JWT token
	// For now, we'll use a placeholder
	userID := uint(1)

	err = h.favoriteService.RemoveFavorite(userID, uint(riddleID))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to remove favorite")
	}

	return c.NoContent(http.StatusNoContent)
}