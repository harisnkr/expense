package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/harisnkr/expense/data"
	"github.com/harisnkr/expense/dto"
	"github.com/harisnkr/expense/models"
	"net/http"
)

func UpdateCard(ctx *gin.Context) {
	// parse request body and path parameter
	var req dto.UpdateCardReq
	ctx.Bind(&req)

	// find card to update using `id` path parameter
	var card models.Card
	data.DB.First(&card, ctx.Param("id"))

	// update
	data.DB.Model(&card).Updates(
		models.Card{
			IssuerBank:     req.IssuerBank,
			Name:           req.Name,
			Network:        req.Network,
			MilesPerDollar: req.MilesPerDollar,
		},
	)

	ctx.JSON(http.StatusOK, gin.H{
		"updated_card": card,
	})
}
