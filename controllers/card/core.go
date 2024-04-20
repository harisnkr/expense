package card

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/harisnkr/expense/common"
	"github.com/harisnkr/expense/data"
	"github.com/harisnkr/expense/models"
)

// API is an interface for operations related to models.Card
type API interface {
	AdminCreateCard(ctx *gin.Context)
	AdminDeleteCard(ctx *gin.Context)
	GetCard(ctx *gin.Context)
	GetAllCards(ctx *gin.Context)
	AdminUpdateCard(ctx *gin.Context)
	AddCardToUser(ctx *gin.Context)
}

// Impl holds dependencies for card.API
type Impl struct {
	database    *mongo.Client
	collections *data.Collections
}

// New returns Impl struct with dependencies for using card.API
func New(database *mongo.Client, collections *data.Collections) *Impl {
	return &Impl{database, collections}
}

// GetCard gets one card by name, issuerBank and network
func (a *Impl) GetCard(c *gin.Context) {
	var (
		cards         []models.Card
		reqName       = c.Param("name")
		reqIssuerBank = c.Param("issuerBank")
		reqNetwork    = c.Param("network")
		log           = slog.With(common.RequestID, c.MustGet(common.RequestID)).
				With("func", "GetCard",
				"reqName", reqName, "reqIssuerBank", reqIssuerBank, "reqNetwork", reqNetwork)
	)

	log.Debug("incoming req to search for card")
	cursor, err := a.collections.Cards.Find(c, bson.D{
		{"name", c.Param("name")},
		{"issuer_bank", reqIssuerBank},
		{"network", reqNetwork},
	})
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

// GetAllCards gets all cards available
func (a *Impl) GetAllCards(c *gin.Context) {
	log := slog.With(common.RequestID, c.MustGet(common.RequestID))
	log.Debug("getting all cards")
	cursor, err := a.collections.Cards.Find(c, bson.M{})
	if err != nil {
		log.Warn("find card error: ", err)
		return
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		if err := cursor.Close(ctx); err != nil {
			log.Error("cursor.Close: ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}(cursor, c)

	var results []models.Card
	if err = cursor.All(c, &results); err != nil {
		log.Warn("GetAllCards failed", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"cards": results,
	})
}
