package card

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/harisnkr/expense/common"
	"github.com/harisnkr/expense/dto"
	"github.com/harisnkr/expense/models"
)

// AdminCreateCard creates a card object
func (a *Impl) AdminCreateCard(c *gin.Context) {
	var (
		card models.Card
		log  = slog.With(common.RequestID, c.MustGet(common.RequestID))
	)
	if err := c.ShouldBindJSON(&card); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	card.ID = uuid.New().String()
	result, err := a.collections.Cards.InsertOne(c, card)
	if err != nil {
		log.Error("Failed to create card", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}

// AdminDeleteCard deletes a card object
func (a *Impl) AdminDeleteCard(c *gin.Context) {
	var (
		req *dto.AdminDeleteCardRequest
		log = slog.With(common.RequestID, c.MustGet(common.RequestID))
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	deleteResult, err := a.collections.Cards.DeleteOne(c,
		bson.M{"name": req.Name, "issuer_bank": req.IssuerBank, "network": req.Network})
	if err != nil {
		log.Warn("delete card error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, deleteResult)
}

// AdminUpdateCard updates a card object
func (a *Impl) AdminUpdateCard(c *gin.Context) {
	var (
		log = slog.With(common.RequestID, c.MustGet(common.RequestID))
		req *dto.AdminUpdateCardRequest
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := a.collections.Cards.UpdateOne(c,
		bson.M{"_id": req.ID},       // cardID to update
		bson.M{"$set": req.Updates}) // fields to update
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if result.ModifiedCount == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "card not updated"})
		return
	}

	var updatedCard models.Card
	if err = a.collections.Cards.FindOne(c, bson.M{"_id": req.ID}).Decode(&updatedCard); err != nil {
		log.Warn("find updated card error: ", err)
	}

	log.Debug("updated card successfully")
	c.JSON(http.StatusOK, updatedCard)
}
