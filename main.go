package main

import (
	"context"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/harisnkr/expense/common"
	"github.com/harisnkr/expense/common/logkeys"
	"github.com/harisnkr/expense/config"
	"github.com/harisnkr/expense/controllers"
	"github.com/harisnkr/expense/controllers/card"
	"github.com/harisnkr/expense/controllers/user"
	"github.com/harisnkr/expense/data"
)

var (
	cardAPI card.API
	userAPI user.API
)

func main() {
	r := gin.Default()

	// TODO: add app config
	client, collections := data.InitDatabase(context.Background())

	cardAPI = card.New(client, collections)
	userAPI = user.New(client, collections)

	// routes
	r.GET("/health", controllers.Health)
	registerCardsRoutes(r, cardAPI)
	registerUsersRoutes(r, userAPI)

	if err := r.Run(); err != nil {
		log.Error(logkeys.Error, err, "Failed to start server")
		return
	}
}

func registerUsersRoutes(r *gin.Engine, userAPI user.API) {
	r.POST("/user/register", userAPI.RegisterUser)
	r.POST("/user/email/verify", userAPI.VerifyEmail)
	r.POST("/user/login", userAPI.LoginUser)
	r.PATCH("/user", userAPI.UpdateUser)
}

func registerCardsRoutes(r *gin.Engine, cardAPI card.API) {
	// TODO: use adminGroup := router.Group("/admin")
	r.POST("/admin/card", cardAPI.CreateCard)
	r.PUT("/admin/card/:id", cardAPI.UpdateCard)
	r.DELETE("/admin/card/:id", cardAPI.DeleteCard)

	// user facing
	r.GET("/cards", cardAPI.SearchCards)
	r.GET("/card/:id", cardAPI.SearchCard)
}

func init() {
	log.SetLevel(log.DebugLevel)
	config.InitEnvVar()
	common.InitValidators()
}
