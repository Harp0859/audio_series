package main

import (
	"log"
	"os"

	"audio-series-app/backend/internal/config"
	"audio-series-app/backend/internal/middleware"

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

	// Initialize middleware
	corsMiddleware := middleware.NewCorsMiddleware(cfg)

	// Create router
	router := gin.Default()

	// Apply middleware
	router.Use(corsMiddleware.Handle())

	// Serve static files (frontend)
	router.Static("/static", "./frontend")
	router.LoadHTMLGlob("frontend/*.html")

	// Serve the main frontend page
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "audio_series_frontend.html", nil)
	})

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":       "ok",
			"message":      "Audio Series App is running",
			"supabase_url": cfg.SupabaseURL,
		})
	})

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3003"
	}

	log.Printf("🚀 Server starting on port %s", port)
	log.Printf("📊 Environment: %s", cfg.Environment)
	log.Printf("🔗 Supabase URL: %s", cfg.SupabaseURL)

	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
