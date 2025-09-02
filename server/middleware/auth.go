package middleware

import (
	"net/http"
	"strings"
	"riddles-server/services"

	"github.com/labstack/echo/v4"
)

type AuthMiddleware struct {
	authService services.AuthService
}

func NewAuthMiddleware(authService services.AuthService) *AuthMiddleware {
	return &AuthMiddleware{
		authService: authService,
	}
}

func (m *AuthMiddleware) AuthRequired(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing authorization header")
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid authorization header format")
		}

		_, err := m.authService.ValidateAccessToken(tokenString)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or expired token")
		}

		// TODO: Extract user ID from token and store in context for use in handlers
		// For now, we'll just validate the token

		return next(c)
	}
}