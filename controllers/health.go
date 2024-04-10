package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Health(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"message": "expense service is up!",
	})
}
