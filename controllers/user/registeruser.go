package user

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"

	"github.com/harisnkr/expense/dto"
	"github.com/harisnkr/expense/models"
)

// RegisterUser registers a new user with username, password, and email
func (u *Impl) RegisterUser(c *gin.Context) {
	var (
		collection = u.collections.Users
		req        *dto.RegisterUserRequest
		log        = slog.With(c)
	)
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if email in request already exists in database
	var existingUser models.User
	err := collection.FindOne(c, bson.M{"email": req.Email}).Decode(&existingUser)
	if err == nil { // if no error (email was found)
		if strings.EqualFold(existingUser.Email, req.Email) {
			c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
			return
		}
	}

	// if not proceed on to create newUser
	newUser := &models.User{}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	otp := populateUserEntry(newUser, req, hashedPassword)

	// Insert the new user into the database
	if _, err = collection.InsertOne(c, newUser); err != nil {
		log.Error("Failed to insert new user", err)
		// TODO: create generic handlers for errors
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Unknown error"})
		return
	}

	// Send email with verification link
	go sendVerificationEmail(c, req.Email, otp)
	c.JSON(http.StatusCreated, gin.H{"message": "Check email for verification code."})
}
