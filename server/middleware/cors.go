package middleware

import (
	"github.com/labstack/echo/v4/middleware"
)

func CORS() middleware.CORSConfig {
	return middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", "http://localhost:3001"}, // Frontend URLs
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}
}