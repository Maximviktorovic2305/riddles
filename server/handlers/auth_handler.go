package handlers

import (
	"net/http"
	"riddles-server/models"
	"riddles-server/services"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

type RegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type AuthResponse struct {
	User         *models.User `json:"user"`
	AccessToken  string       `json:"access_token"`
	RefreshToken string       `json:"refresh_token"`
}

func (h *AuthHandler) Register(c echo.Context) error {
	var req RegisterRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	user, err := h.authService.Register(req.Username, req.Email, req.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Generate tokens
	accessToken, refreshToken, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to generate tokens")
	}

	return c.JSON(http.StatusCreated, AuthResponse{
		User:         user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

func (h *AuthHandler) Login(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	user, accessToken, refreshToken, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
	}

	return c.JSON(http.StatusOK, AuthResponse{
		User:         user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type RefreshResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (h *AuthHandler) Refresh(c echo.Context) error {
	var req RefreshRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	accessToken, refreshToken, err := h.authService.RefreshToken(req.RefreshToken)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid refresh token")
	}

	return c.JSON(http.StatusOK, RefreshResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}