package main

import (
	"log"
	"riddles-server/database"
	"riddles-server/routes"
	"riddles-server/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Initialize database
	database.ConnectDB()
	database.MigrateDB()

	// Select daily riddles for today
	if err := utils.SelectDailyRiddles(); err != nil {
		log.Printf("Warning: Failed to select daily riddles: %v", err)
	} else {
		log.Println("Daily riddles selected successfully")
	}

	// Create Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", "http://localhost:3001", "http://localhost:3002", "http://localhost:3003"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	// Setup routes
	routes.SetupRoutes(e)

	// Start server
	log.Println("Starting server on port 8080...")
	e.Logger.Fatal(e.Start(":8080"))
}