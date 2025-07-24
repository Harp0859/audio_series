package main

import (
	"log"
	"os"

	"audio-series-app/backend/internal/config"
	"audio-series-app/backend/internal/handlers"
	"audio-series-app/backend/internal/middleware"
	"audio-series-app/backend/internal/routes"
	"audio-series-app/backend/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Initialize configuration
	cfg := config.Load()

	// Set Gin mode
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize services
	supabaseService := services.NewSupabaseService(cfg)
	authService := services.NewAuthService(cfg, supabaseService)
	userService := services.NewUserService(supabaseService)
	seriesService := services.NewSeriesService(supabaseService)
	episodeService := services.NewEpisodeService(supabaseService)
	paymentService := services.NewPaymentService(cfg, supabaseService)
	coinService := services.NewCoinService(supabaseService)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)
	seriesHandler := handlers.NewSeriesHandler(seriesService)
	episodeHandler := handlers.NewEpisodeHandler(episodeService, coinService)
	paymentHandler := handlers.NewPaymentHandler(paymentService, coinService)
	adminHandler := handlers.NewAdminHandler(seriesService, episodeService, userService)

	// Initialize middleware
	authMiddleware := middleware.NewAuthMiddleware(authService)
	corsMiddleware := middleware.NewCorsMiddleware(cfg)

	// Create router
	router := gin.Default()

	// Apply middleware
	router.Use(corsMiddleware.Handle())

	// Setup routes
	routes.SetupRoutes(
		router,
		authHandler,
		userHandler,
		seriesHandler,
		episodeHandler,
		paymentHandler,
		adminHandler,
		authMiddleware,
	)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3003"
	}

	log.Printf("🚀 Server starting on port %s", port)
	log.Printf("📊 Environment: %s", cfg.Environment)

	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
