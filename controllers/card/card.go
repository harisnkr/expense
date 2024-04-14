package card

import (
	slog "log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/harisnkr/expense/data"
	"github.com/harisnkr/expense/dto"
	"github.com/harisnkr/expense/models"
)

// API is an interface for operations related to models.Card
type API interface {
	AdminCreateCard(ctx *gin.Context)
	AdminDeleteCard(ctx *gin.Context)
	SearchCard(ctx *gin.Context)
	SearchCards(ctx *gin.Context)
	AdminUpdateCard(ctx *gin.Context)
	AddCardToUser(ctx *gin.Context)
}

type Impl struct {
	database    *mongo.Client
	collections *data.Collections
}

func New(database *mongo.Client, collections *data.Collections) *Impl {
	return &Impl{database, collections}
}

// AdminCreateCard creates a card object
func (a *Impl) AdminCreateCard(c *gin.Context) {
	var (
		card models.Card
		log  = slog.With(c)
	)
	if err := c.ShouldBindJSON(&card); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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
		log = slog.With(c)
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

// SearchCard search for one card available
func (a *Impl) SearchCard(c *gin.Context) {
	var (
		cards []models.Card
		log   = slog.With(c)
	)

	log.Debug("incoming req to search for card with name: " + c.Param("name"))

	cursor, err := a.collections.Cards.Find(c, bson.D{{"name", c.Param("name")}})
	if err != nil {
		log.Warn("find card error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err = cursor.All(c, &cards); err != nil {
		log.Warn("cursor.All error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(cards) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "no card result found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": cards,
	})
}

// SearchCards searches for all cards available
func (a *Impl) SearchCards(c *gin.Context) {
	var (
		cards []models.Card
		log   = slog.With(c)
	)

	cursor, err := a.collections.Cards.Find(c, nil) // filter = nil to search all
	if err != nil {
		return
	}
	defer cursor.Close(c)

	// Slice to store the decoded documents
	var results []models.Card

	// Decode all documents into the slice
	if err = cursor.All(c, &results); err != nil {
		log.Warn("SearchCards failed", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"cards": cards,
	})
}

func (a *Impl) AdminUpdateCard(c *gin.Context) {
	var (
		log = slog.With(c)
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

func (a *Impl) AddCardToUser(c *gin.Context) {
	var (
		userID = c.GetString("userID") // get from userID set from JWT auth
		log    = slog.With(c, "func", "AddCardToUser", "userID", userID)
	)

	var req *dto.AddCardToUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log = log.With("cardID", req.CardID)

	var card models.Card
	if err := a.collections.Cards.FindOne(c, bson.M{"card_id": req.CardID}).Decode(&card); err != nil {
		log.Error("Failed to fetch card", err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	log.Debug("found card to add",
		"card_id", card,
		"name", card.Name,
		"issuer_bank", card.IssuerBank,
		"network", card.Network,
	)

	result, err := a.collections.Users.UpdateOne(c,
		bson.M{"_id": userID},
		bson.M{"$push": bson.M{"cards": card}},
	)
	if err != nil {
		log.Error("Failed to add card to user", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if result.MatchedCount == 0 {
		log.Error("Failed to add card to user", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to add card to user"})
		return
	}
	log.Debug("successfully add card to user")
}
