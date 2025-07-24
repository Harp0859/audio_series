package routes

import (
	"audio-series-app/backend/internal/handlers"
	"audio-series-app/backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	router *gin.Engine,
	authHandler *handlers.AuthHandler,
	userHandler *handlers.UserHandler,
	seriesHandler *handlers.SeriesHandler,
	episodeHandler *handlers.EpisodeHandler,
	paymentHandler *handlers.PaymentHandler,
	adminHandler *handlers.AdminHandler,
	authMiddleware *middleware.AuthMiddleware,
) {
	// Root route
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Audio Series API",
			"version": "1.0.0",
			"status":  "running",
			"docs":    "/api/v1/health",
		})
	})

	// API version
	api := router.Group("/api/v1")

	// Health check
	api.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "message": "Audio Series API is running"})
	})

	// Public routes
	public := api.Group("/")
	{
		// Authentication
		public.POST("/auth/register", authHandler.Register)
		public.POST("/auth/login", authHandler.Login)
		public.POST("/auth/refresh", authHandler.RefreshToken)

		// Series (public read access)
		public.GET("/series", seriesHandler.GetSeries)
		public.GET("/series/:id", seriesHandler.GetSeriesByID)

		// Payment bundles
		public.GET("/payment/bundles", paymentHandler.GetCoinBundles)
	}

	// Protected routes (require authentication)
	protected := api.Group("/")
	protected.Use(authMiddleware.Authenticate())
	{
		// User routes
		protected.GET("/user/profile", userHandler.GetProfile)
		protected.GET("/user/purchases", userHandler.GetPurchases)
		protected.GET("/user/coins", userHandler.GetCoinBalance)

		// Episodes
		protected.GET("/episodes/:id", episodeHandler.GetEpisode)
		protected.POST("/episodes/:id/unlock", episodeHandler.UnlockEpisode)
		protected.POST("/series/:id/unlock", episodeHandler.UnlockSeries)

		// Payments
		protected.POST("/payment/initiate", paymentHandler.InitiatePayment)
	}

	// Admin routes (require admin role)
	admin := api.Group("/admin")
	admin.Use(authMiddleware.Authenticate())
	admin.Use(authMiddleware.RequireAdmin())
	{
		admin.POST("/series", adminHandler.CreateSeries)
		admin.POST("/episodes", adminHandler.CreateEpisode)
		admin.GET("/stats", adminHandler.GetAdminStats)
	}

	// Payment callbacks (public)
	api.POST("/payment/callback/:gateway", paymentHandler.PaymentCallback)
}
