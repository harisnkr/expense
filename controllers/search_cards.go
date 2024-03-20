package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/harisnkr/expense/data"
	"github.com/harisnkr/expense/models"

	"net/http"
)

func SearchCards(ctx *gin.Context) {
	var (
		cards []models.Card
	)
	data.DB.Find(&cards)

	ctx.JSON(http.StatusOK, gin.H{
		"cards": cards,
	})
}
