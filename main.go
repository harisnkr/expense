package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/harisnkr/expense/common/logkeys"
	"github.com/harisnkr/expense/config"
	"github.com/harisnkr/expense/controllers"
	"github.com/harisnkr/expense/controllers/cards"
	"github.com/harisnkr/expense/data"
)

var cardsAPI cards.API

func main() {
	r := gin.Default()

	db := data.InitDatabase()

	// implementations
	cardsAPI = cards.New(db)

	// routes
	r.GET("/health", controllers.Health)
	registerCardsRoutes(r, cardsAPI)

	if err := r.Run(); err != nil {
		log.Error(logkeys.Error, err, "Failed to start server")
		return
	}
}

func registerCardsRoutes(r *gin.Engine, cardsAPI cards.API) {
	r.POST("/card", cardsAPI.CreateCard)
	r.GET("/cards", cardsAPI.SearchCards)
	r.GET("/card/:id", cardsAPI.SearchCard)
	r.PUT("/card/:id", cardsAPI.UpdateCard)
	r.DELETE("/card/:id", cardsAPI.DeleteCard)
}

func init() {
	log.SetLevel(log.DebugLevel)
	config.InitEnvVar()

	// TODO: add redis for fallback and speed
}
