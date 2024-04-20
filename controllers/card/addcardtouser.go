package card

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/harisnkr/expense/dto"
	"github.com/harisnkr/expense/models"
)

// AddCardToUser adds a card to user
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
	if err := a.collections.Cards.FindOne(c, bson.M{"_id": req.CardID}).Decode(&card); err != nil {
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
