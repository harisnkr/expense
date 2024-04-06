package card

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/harisnkr/expense/data"
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
	var card models.Card
	if err := c.ShouldBindJSON(&card); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := a.collections.Cards.InsertOne(c, card)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusCreated, result)
}

// AdminDeleteCard deletes a card object
func (a *Impl) AdminDeleteCard(c *gin.Context) {
	//var card models.Cards
	//a.database.Delete(&card, c.Param("id"))
	c.Status(http.StatusOK)
}

// SearchCard search for one card available TODO: make this search by `Name` from models.Card
func (a *Impl) SearchCard(c *gin.Context) {
	//var card models.Cards
	//a.database.First(&card, c.Param("id"))

	//c.JSON(http.StatusOK, gin.H{
	//	"card": card,
	//})
}

// SearchCards searches for all cards available
func (a *Impl) SearchCards(c *gin.Context) {
	var (
		cards []models.Card
	)
	//a.database.Find(&cards)

	c.JSON(http.StatusOK, gin.H{
		"cards": cards,
	})
}

func (a *Impl) AdminUpdateCard(c *gin.Context) {
	// parse request body and path parameter
	//var req dto.UpdateCardReq
	//c.Bind(&req)

	// find card to update using `id` path parameter
	var card models.Card
	//a.database.First(&card, c.Param("id"))

	// update
	//a.database.Model(&card).Updates(
	//	models.Cards{
	//		IssuerBank: req.IssuerBank,
	//		Name:       req.Name,
	//		Network:    req.Network,
	//		Miles:      req.Miles,
	//	},
	//)

	c.JSON(http.StatusOK, gin.H{
		"updated_card": card,
	})
}

func (a *Impl) AddCardToUser(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
