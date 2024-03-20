package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/harisnkr/expense/data"
	"github.com/harisnkr/expense/models"

	"net/http"
)

func SearchCard(ctx *gin.Context) {
	var card models.Card

	data.DB.First(&card, ctx.Param("id"))

	ctx.JSON(http.StatusOK, gin.H{
		"card": card,
	})
}
