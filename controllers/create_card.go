package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/harisnkr/expense/common/logkeys"
	"github.com/harisnkr/expense/data"
	"github.com/harisnkr/expense/dto"
	"github.com/harisnkr/expense/models"

	log "github.com/sirupsen/logrus"
	"net/http"
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
