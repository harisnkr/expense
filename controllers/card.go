package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harisnkr/expense/common/logkeys"
	"github.com/harisnkr/expense/data"
	"github.com/harisnkr/expense/dto"
	"github.com/harisnkr/expense/models"

	log "github.com/sirupsen/logrus"
)

// CreateCard creates a card object
func CreateCard(ctx *gin.Context) {
	req := dto.CreateCardReq{}
	ctx.Bind(&req)

	card := models.Card{
		IssuerBank:     req.IssuerBank,
		Name:           req.Name,
		Network:        req.Network,
		MilesPerDollar: req.MilesPerDollar,
	}
	result := data.DB.Create(&card)
	if result.Error != nil {
		log.Error(logkeys.Error, result.Error, "something went wrong when creating card")
	}

	log.Info("successfully stored card details!")

	ctx.JSON(http.StatusOK, gin.H{
		"card": card,
	})
}

func DeleteCard(ctx *gin.Context) {
	var card models.Card
	data.DB.Delete(&card, ctx.Param("id"))
	ctx.Status(http.StatusOK)
}

func SearchCard(ctx *gin.Context) {
	var card models.Card

	data.DB.First(&card, ctx.Param("id"))

	ctx.JSON(http.StatusOK, gin.H{
		"card": card,
	})
}

func SearchCards(ctx *gin.Context) {
	var (
		cards []models.Card
	)
	data.DB.Find(&cards)

	ctx.JSON(http.StatusOK, gin.H{
		"cards": cards,
	})
}

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
