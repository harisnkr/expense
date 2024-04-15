package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Health is a simple health check endpoint
func Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "expense service is up!"})
}
