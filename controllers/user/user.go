package user

import (
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/harisnkr/expense/data"
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
		user       models.User
	)
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the username or email already exists
	var existingUser models.User
	err := collection.FindOne(c, bson.M{"$or": []bson.M{{"username": user.Username}, {"email": user.Email}}}).Decode(&existingUser)
	if err == nil {
		if strings.EqualFold(existingUser.Username, user.Username) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		}
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)
	otp := generateToken()

	user.VerificationCode = otp
	user.VerificationSentAt = time.Now()
	user.UpdatedAt = time.Now()
	user.CreatedAt = time.Now()
	user.ID = uuid.New().String()

	// Insert the new user into the database
	if _, err = collection.InsertOne(c, user); err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Unknown error"})
	}

	// Send email with verification link
	go sendVerificationEmail(user.Email, otp)

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

// generateToken generates a random verification token
func generateToken() string {
	// Generate a random 8-digit number
	otp := rand.Intn(100000000)

	// Ensure the OTP is exactly 8 digits long
	return fmt.Sprintf("%08d", otp)
}

// sendVerificationEmail sends an email with the verification link
func sendVerificationEmail(email, token string) {
	// Implement your email sending logic here
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
