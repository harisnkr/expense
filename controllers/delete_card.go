package controllers

import (
	"github.com/harisnkr/expense/data"
	"github.com/harisnkr/expense/models"

	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteCard(ctx *gin.Context) {
	var card models.Card
	data.DB.Delete(&card, ctx.Param("id"))
	ctx.Status(http.StatusOK)
}
