package main

import (
	"github.com/gin-gonic/gin"
	"github.com/harisnkr/expense/common/logkeys"
	"github.com/harisnkr/expense/config"
	"github.com/harisnkr/expense/controllers"
	"github.com/harisnkr/expense/data"

	log "github.com/sirupsen/logrus"
)

func main() {
	r := gin.Default()
	r.GET("/health", controllers.Health)

	// TODO: segregate routes. e.g. cards, savings accounts, expenditure, items, etc.
	r.POST("/card", controllers.CreateCard)
	r.GET("/cards", controllers.SearchCards)
	r.GET("/card/:id", controllers.SearchCard)
	r.PUT("/card/:id", controllers.UpdateCard)
	r.DELETE("/card/:id", controllers.DeleteCard)

	if err := r.Run(); err != nil {
		log.Error(logkeys.Error, err, "Failed to start server")
		return
	}
}

func init() {
	log.SetLevel(log.DebugLevel)
	config.InitEnvVar()
	// TODO: dependency injection for database
	data.InitDatabase()

	// TODO: add redis for fallback and speed
}
