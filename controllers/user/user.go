package user

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"

	"github.com/harisnkr/expense/common"
	"github.com/harisnkr/expense/data"
	"github.com/harisnkr/expense/dto"
	"github.com/harisnkr/expense/models"
)

// API is an interface for operations related to models.User
type API interface {
	RegisterUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	VerifyEmail(ctx *gin.Context)
	LoginUser(ctx *gin.Context)
}

type Impl struct {
	database    *mongo.Client
	collections *data.Collections
}

func New(database *mongo.Client, collections *data.Collections) *Impl {
	return &Impl{database, collections}
}

// RegisterUser registers a new user with username, password, and email
func (u *Impl) RegisterUser(c *gin.Context) {
	var (
		collection = u.collections.Users
		req        *dto.RegisterUserRequest
	)
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the username or email already exists
	var existingUser models.User
	err := collection.FindOne(c, bson.M{"email": req.Email}).Decode(&existingUser)
	if err == nil {
		if strings.EqualFold(existingUser.Email, req.Email) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		}
		return
	}

	// if not proceed on to create newUser
	var newUser models.User
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	newUser.Password = string(hashedPassword)
	otp := common.GenerateToken()

	newUser.VerificationCode = otp
	newUser.VerificationSentAt = time.Now()
	newUser.UpdatedAt = time.Now()
	newUser.CreatedAt = time.Now()
	newUser.ID = uuid.New().String()

	// Insert the new req into the database
	if _, err = collection.InsertOne(c, req); err != nil {
		log.Fatal(err)
		// TODO: create generic handlers for errors
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Unknown error"})
	}

	// Send email with verification link
	go sendVerificationEmail(req.Email, otp)
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully. Check email for verification."})
}

// VerifyEmail verifies the email with the verification token
func (u *Impl) VerifyEmail(c *gin.Context) {
	var (
		email      = c.Query("email")
		token      = c.Query("token")
		collection = u.collections.Users
	)

	// Fetch the user by email and verification token
	var user models.User
	err := u.collections.Users.FindOne(c, bson.M{"email": email, "verification_code": token}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid verification code"})
		return
	}

	// Mark the user as verified
	update := bson.M{"$set": bson.M{"verified": true}}
	if _, err = collection.UpdateOne(c, bson.M{"email": email}, update); err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Unknown error"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email verified successfully"})
}

// sendVerificationEmail sends an email with the verification link
func sendVerificationEmail(email, token string) {
	// TODO: Implement email sending logic here
	log.Debugf("Sending verification email to %s with OTP: %s", email, token)
}

func (u *Impl) UpdateUser(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (u *Impl) DeleteUser(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (u *Impl) LoginUser(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
