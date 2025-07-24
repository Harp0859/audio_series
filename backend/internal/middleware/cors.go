package middleware

import (
	"audio-series-app/backend/internal/config"

	"github.com/gin-gonic/gin"
)

type CorsMiddleware struct {
	config *config.Config
}

func NewCorsMiddleware(cfg *config.Config) *CorsMiddleware {
	return &CorsMiddleware{
		config: cfg,
	}
}

func (m *CorsMiddleware) Handle() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Header("Access-Control-Expose-Headers", "Content-Length")
		c.Header("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})
}
