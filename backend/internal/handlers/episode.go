package handlers

import (
	"net/http"

	"audio-series-app/backend/internal/services"

	"github.com/gin-gonic/gin"
)

type EpisodeHandler struct {
	episodeService *services.EpisodeService
	coinService    *services.CoinService
}

func NewEpisodeHandler(episodeService *services.EpisodeService, coinService *services.CoinService) *EpisodeHandler {
	return &EpisodeHandler{
		episodeService: episodeService,
		coinService:    coinService,
	}
}

// GetEpisode returns a specific episode with purchase status
func (h *EpisodeHandler) GetEpisode(c *gin.Context) {
	episodeIDStr := c.Param("id")
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userIDStr, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}

	// Parse UUIDs (simplified for now)
	episodeWithPurchase, err := h.episodeService.GetEpisodeWithPurchaseStatus(c.Request.Context(), episodeIDStr, userIDStr)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Episode not found"})
		return
	}

	c.JSON(http.StatusOK, episodeWithPurchase)
}

// UnlockEpisode unlocks an episode using coins
func (h *EpisodeHandler) UnlockEpisode(c *gin.Context) {
	episodeIDStr := c.Param("id")
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userIDStr, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}

	err := h.coinService.UnlockEpisode(c.Request.Context(), userIDStr, episodeIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Episode unlocked successfully"})
}

// UnlockSeries unlocks an entire series using coins
func (h *EpisodeHandler) UnlockSeries(c *gin.Context) {
	seriesIDStr := c.Param("id")
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userIDStr, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}

	err := h.coinService.UnlockSeries(c.Request.Context(), userIDStr, seriesIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Series unlocked successfully"})
}
