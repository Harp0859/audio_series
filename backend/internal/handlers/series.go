package handlers

import (
	"net/http"

	"audio-series-app/backend/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SeriesHandler struct {
	seriesService *services.SeriesService
}

func NewSeriesHandler(seriesService *services.SeriesService) *SeriesHandler {
	return &SeriesHandler{
		seriesService: seriesService,
	}
}

// GetSeries returns all available series
func (h *SeriesHandler) GetSeries(c *gin.Context) {
	series, err := h.seriesService.GetSeries(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get series"})
		return
	}

	c.JSON(http.StatusOK, series)
}

// GetSeriesByID returns a specific series with its episodes
func (h *SeriesHandler) GetSeriesByID(c *gin.Context) {
	seriesIDStr := c.Param("id")
	seriesID, err := uuid.Parse(seriesIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid series ID"})
		return
	}

	seriesWithEpisodes, err := h.seriesService.GetSeriesWithEpisodes(c.Request.Context(), seriesID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Series not found"})
		return
	}

	c.JSON(http.StatusOK, seriesWithEpisodes)
}
