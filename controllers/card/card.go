package card

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
		log  = logrus.WithContext(c)
	)
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
	var (
		req *dto.AdminDeleteCardRequest
		log = logrus.WithContext(c)
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	deleteResult, err := a.collections.Cards.DeleteOne(c, bson.M{"name": req.Name, "issuer_bank": req.IssuerBank})
	if err != nil {
		log.Warn("delete card error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, deleteResult)
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
		log   = logrus.WithContext(c)
	)

	cursor, err := a.collections.Cards.Find(c, nil)
	if err != nil {
		return
	}
	defer cursor.Close(c)

	// Slice to store the decoded documents
	var results []models.Card

	// Decode all documents into the slice
	if err = cursor.All(c, &results); err != nil {
		log.Warn("SearchCards failed with err: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"cards": cards,
	})
}

func (a *Impl) AdminUpdateCard(c *gin.Context) {
	//parse request body and path parameter
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
