package handlers

import (
	"net/http"

	"audio-series-app/backend/internal/models"
	"audio-series-app/backend/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AdminHandler struct {
	seriesService  *services.SeriesService
	episodeService *services.EpisodeService
	userService    *services.UserService
}

func NewAdminHandler(seriesService *services.SeriesService, episodeService *services.EpisodeService, userService *services.UserService) *AdminHandler {
	return &AdminHandler{
		seriesService:  seriesService,
		episodeService: episodeService,
		userService:    userService,
	}
}

// CreateSeries creates a new series (admin only)
func (h *AdminHandler) CreateSeries(c *gin.Context) {
	var series models.Series
	if err := c.ShouldBindJSON(&series); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Parse user ID
	userIDStr, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}

	userUUID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	// Set the creator
	series.CreatedBy = userUUID

	err = h.seriesService.CreateSeries(c.Request.Context(), &series)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create series"})
		return
	}

	c.JSON(http.StatusCreated, series)
}

// CreateEpisode creates a new episode (admin only)
func (h *AdminHandler) CreateEpisode(c *gin.Context) {
	var episode models.Episode
	if err := c.ShouldBindJSON(&episode); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	err := h.episodeService.CreateEpisode(c.Request.Context(), &episode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create episode"})
		return
	}

	c.JSON(http.StatusCreated, episode)
}

// GetAdminStats returns admin dashboard statistics
func (h *AdminHandler) GetAdminStats(c *gin.Context) {
	// This would typically get stats from the database
	// For now, we'll return mock data
	stats := models.AdminStats{
		TotalUsers:     100,
		TotalSeries:    5,
		TotalEpisodes:  25,
		TotalRevenue:   50000,
		MonthlyRevenue: 10000,
		ActiveUsers:    75,
	}

	c.JSON(http.StatusOK, stats)
}
