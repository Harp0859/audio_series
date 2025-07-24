package handlers

import (
	"net/http"

	"audio-series-app/backend/internal/models"
	"audio-series-app/backend/internal/services"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	paymentService *services.PaymentService
	coinService    *services.CoinService
}

func NewPaymentHandler(paymentService *services.PaymentService, coinService *services.CoinService) *PaymentHandler {
	return &PaymentHandler{
		paymentService: paymentService,
		coinService:    coinService,
	}
}

// GetCoinBundles returns available coin bundles for purchase
func (h *PaymentHandler) GetCoinBundles(c *gin.Context) {
	currency := c.DefaultQuery("currency", "INR")
	bundles := h.paymentService.GetCoinBundles(currency)
	c.JSON(http.StatusOK, bundles)
}

// InitiatePayment starts a payment process
func (h *PaymentHandler) InitiatePayment(c *gin.Context) {
	var req models.PaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

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

	response, err := h.paymentService.InitiatePayment(c.Request.Context(), userIDStr, req.BundleID, req.Currency)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// PaymentCallback handles payment gateway callbacks
func (h *PaymentHandler) PaymentCallback(c *gin.Context) {
	gateway := c.Param("gateway")

	var paymentData map[string]interface{}
	if err := c.ShouldBindJSON(&paymentData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment data"})
		return
	}

	err := h.paymentService.HandlePaymentCallback(c.Request.Context(), gateway, paymentData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment processed successfully"})
}
