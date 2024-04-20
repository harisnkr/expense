package user

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"

	"github.com/harisnkr/expense/common"
	"github.com/harisnkr/expense/dto"
	"github.com/harisnkr/expense/models"
)

// Login logs in the user with username and password (TODO: google login integration)
func (u *Impl) Login(c *gin.Context) {
	var (
		req dto.UserLoginRequest
		log = slog.With(common.RequestID, c.MustGet(common.RequestID))
	)
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := u.collections.Users.FindOne(c, bson.M{"email": req.Email}).Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Info("User not found")
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		log.Info("Failed to find user", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log = log.With("email", user.Email)
	if !user.Verified {
		log.Info("User is not verified")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User is not verified"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(req.Password), []byte(user.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		log.Warn("Invalid password entered for user", err)
		return
	}

	tokenDuration, tokenString := generateSessionJWT(c, user)
	if tokenString == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to generate session token"})
		return
	}
	c.JSON(http.StatusOK, dto.UserLoginResponse{
		SessionToken: tokenString,
		ExpiresIn:    tokenDuration.String(),
	})
}
