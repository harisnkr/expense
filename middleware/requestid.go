package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RequestID adds a unique request ID to each incoming request
func RequestID(c *gin.Context) {
	requestID := uuid.New().String()
	requestID = strings.ReplaceAll(requestID, "-", "")

	c.Set("request_id", requestID)
	c.Writer.Header().Set("X-Request-ID", requestID)
	c.Next()
}
