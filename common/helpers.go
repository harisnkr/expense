package common

import (
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Parse binds JSON request body to a struct and handles invalid parameters
func Parse(ctx *gin.Context, req interface{}) error {
	if err := ctx.ShouldBindJSON(req); err != nil {
		log.Warn(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameters"})
		return err
	}
	return nil
}
