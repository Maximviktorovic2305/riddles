package routes

import (
	"riddles-server/handlers"
	"riddles-server/middleware"
	"riddles-server/repository"
	"riddles-server/services"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	// Initialize repositories
	userRepo := repository.NewUserRepository()
	riddleRepo := repository.NewRiddleRepository()
	progressRepo := repository.NewProgressRepository()
	favoriteRepo := repository.NewFavoriteRepository()
	ratingRepo := repository.NewRatingRepository()
	dailyRiddleRepo := repository.NewDailyRiddleRepository()

	// Initialize services
	authService := services.NewAuthService(userRepo, "riddles_secret_key") // In production, use config
	userService := services.NewUserService(userRepo, progressRepo)
	riddleService := services.NewRiddleService(riddleRepo, progressRepo, favoriteRepo, ratingRepo)
	favoriteService := services.NewFavoriteService(favoriteRepo, riddleRepo)
	ratingService := services.NewRatingService(ratingRepo, riddleRepo)
	dailyRiddleService := services.NewDailyRiddleService(dailyRiddleRepo)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)
	riddleHandler := handlers.NewRiddleHandler(riddleService)
	favoriteHandler := handlers.NewFavoriteHandler(favoriteService)
	ratingHandler := handlers.NewRatingHandler(ratingService)
	dailyRiddleHandler := handlers.NewDailyRiddleHandler(dailyRiddleService)

	// Initialize middleware
	authMiddleware := middleware.NewAuthMiddleware(authService)

	// Public routes
	auth := e.Group("/api/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
		auth.POST("/refresh", authHandler.Refresh)
	}

	riddles := e.Group("/api/riddles")
	{
		riddles.GET("", riddleHandler.GetAllRiddles)
		riddles.GET("/:id", riddleHandler.GetRiddleByID)
		riddles.POST("/:id/answer", riddleHandler.CheckAnswer)
	}

	dailyRiddle := e.Group("/api/daily-riddle")
	{
		dailyRiddle.GET("", dailyRiddleHandler.GetTodayRiddle)
		dailyRiddle.GET("/:date", dailyRiddleHandler.GetRiddleByDate)
	}

	// Protected routes
	protected := e.Group("/api")
	protected.Use(authMiddleware.AuthRequired)

	// User routes
	protected.GET("/users/profile", userHandler.GetProfile)
	protected.GET("/users/stats", userHandler.GetUserStats)

	// Favorite routes
	protected.POST("/favorites/:riddle_id", favoriteHandler.AddFavorite)
	protected.DELETE("/favorites/:riddle_id", favoriteHandler.RemoveFavorite)

	// Rating routes
	protected.POST("/ratings/:riddle_id", ratingHandler.RateRiddle)
	protected.DELETE("/ratings/:riddle_id", ratingHandler.RemoveRating)
}