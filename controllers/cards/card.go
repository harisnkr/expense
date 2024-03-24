package cards

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/harisnkr/expense/common/logkeys"
	"github.com/harisnkr/expense/dto"
	"github.com/harisnkr/expense/models"
)

// TODO: Bind -> ShouldBindJSON for all routes for input validation

// API is an interface for operations related to models.Card
type API interface {
	CreateCard(ctx *gin.Context)
	DeleteCard(ctx *gin.Context)
	SearchCard(ctx *gin.Context)
	SearchCards(ctx *gin.Context)
	UpdateCard(ctx *gin.Context)
}

type Impl struct {
	database *gorm.DB
}

func New(database *gorm.DB) *Impl {
	return &Impl{database}
}

// CreateCard creates a card object
func (a *Impl) CreateCard(ctx *gin.Context) {
	req := dto.CreateCardReq{}
	ctx.Bind(&req)

	card := models.Card{
		IssuerBank:     req.IssuerBank,
		Name:           req.Name,
		Network:        req.Network,
		MilesPerDollar: req.MilesPerDollar,
	}
	result := a.database.Create(&card)
	if result.Error != nil {
		log.Error(logkeys.Error, result.Error, "something went wrong when creating card")
	}

	log.Info("successfully stored card details!")

	ctx.JSON(http.StatusOK, gin.H{
		"card": card,
	})
}

func (a *Impl) DeleteCard(ctx *gin.Context) {
	var card models.Card
	a.database.Delete(&card, ctx.Param("id"))
	ctx.Status(http.StatusOK)
}

func (a *Impl) SearchCard(ctx *gin.Context) {
	var card models.Card

	a.database.First(&card, ctx.Param("id"))

	ctx.JSON(http.StatusOK, gin.H{
		"card": card,
	})
}

func (a *Impl) SearchCards(ctx *gin.Context) {
	var (
		cards []models.Card
	)
	a.database.Find(&cards)

	ctx.JSON(http.StatusOK, gin.H{
		"cards": cards,
	})
}

func (a *Impl) UpdateCard(ctx *gin.Context) {
	// parse request body and path parameter
	var req dto.UpdateCardReq
	ctx.Bind(&req)

	// find card to update using `id` path parameter
	var card models.Card
	a.database.First(&card, ctx.Param("id"))

	// update
	a.database.Model(&card).Updates(
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
