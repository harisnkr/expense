package main

import (
	"context"
	log "log/slog"

	"github.com/gin-gonic/gin"

	"github.com/harisnkr/expense/common"
	"github.com/harisnkr/expense/controllers"
	"github.com/harisnkr/expense/controllers/card"
	"github.com/harisnkr/expense/controllers/user"
	"github.com/harisnkr/expense/data"
	"github.com/harisnkr/expense/middleware"
)

var (
	cardAPI card.API
	userAPI user.API
)

func main() {
	r := gin.Default()

	// TODO: add app config?
	client, collections := data.InitDatabase(context.Background())

	cardAPI = card.New(client, collections)
	userAPI = user.New(client, collections)

	r.GET("/health", controllers.Health)

	registerCardRoutes(r, cardAPI)
	registerUserRoutes(r, userAPI)

	if err := r.Run(); err != nil {
		log.Error("Failed to start server, %w", err)
		return
	}
}

func registerUserRoutes(r *gin.Engine, userAPI user.API) {
	userRouter := r.Group("/user")
	{
		userRouter.POST("/register", userAPI.RegisterUser)
		userRouter.POST("/email/verify", userAPI.VerifyEmail)
		userRouter.PATCH("/profile", middleware.Auth(), userAPI.UpdateMe)
	}
	adminRouter := r.Group("/admin")
	{
		adminRouter.GET("user/otp", userAPI.GetEmailOTP)
	}
}

func registerCardRoutes(r *gin.Engine, cardAPI card.API) {
	adminRouter := r.Group("/admin")
	{
		adminRouter.POST("/card", cardAPI.AdminCreateCard)
		adminRouter.PUT("/card/:id", cardAPI.AdminUpdateCard)
		adminRouter.DELETE("/card", cardAPI.AdminDeleteCard)
	}

	r.GET("/cards", cardAPI.SearchCards)
	r.GET("/card/:name", cardAPI.SearchCard)

	r.POST("/user/card/", middleware.Auth(), cardAPI.AddCardToUser)
}

func init() {
	common.SetDependencies()
}
