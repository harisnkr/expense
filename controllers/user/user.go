package user

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"

	"github.com/harisnkr/expense/common"
	"github.com/harisnkr/expense/config"
	"github.com/harisnkr/expense/data"
	"github.com/harisnkr/expense/dto"
	"github.com/harisnkr/expense/models"
)

// API is an interface for operations related to models.User
type API interface {
	RegisterUser(ctx *gin.Context)
	VerifyEmail(ctx *gin.Context)
	UpdateMe(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
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
	newUser := &models.User{}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	otp := populateUserEntry(newUser, req, hashedPassword)

	// Insert the new req into the database
	if _, err = collection.InsertOne(c, newUser); err != nil {
		log.Fatal(err)
		// TODO: create generic handlers for errors
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Unknown error"})
	}

	// Send email with verification link
	go sendVerificationEmail(req.Email, otp)
	c.JSON(http.StatusCreated, gin.H{"message": "Check email for verification code."})
}

// VerifyEmail verifies the email with the verification token
func (u *Impl) VerifyEmail(c *gin.Context) {
	var (
		collection = u.collections.Users
	)

	var req *dto.UserEmailVerifyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Fetch the user by email and verificationCode
	var user models.User
	err := u.collections.Users.FindOne(c, bson.M{"email": req.Email, "verification_code": req.VerificationCode}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid verification code"})
		return
	}

	// Mark the user as verified
	update := bson.M{"$set": bson.M{"verified": true}}
	if _, err = collection.UpdateOne(c, bson.M{"email": user.Email}, update); err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Unknown error"})
	}

	tokenDuration, tokenString := generateSessionJWT(c, user, err)
	c.JSON(http.StatusOK, dto.UserLoginResponse{
		SessionToken: tokenString,
		ExpiresIn:    tokenDuration.String(),
	})
}

func generateSessionJWT(c *gin.Context, user models.User, err error) (time.Duration, string) {
	tokenDuration := time.Hour * 24
	exp := time.Now().Add(tokenDuration).Unix()
	iat := time.Now().Unix()
	nbf := time.Now().Unix()
	iss := common.Issuer
	sub := user.ID
	aud := user.ID

	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"email": user.Email,
		"exp":   exp,
		"iat":   iat,
		"iss":   iss,
		"nbf":   nbf,
		"sub":   sub,
		"aud":   aud,
	})
	tokenString, err := token.SignedString(config.ECDSAKey)
	if err != nil {
		log.Fatal("failed to generate jwt", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Unknown error"})
	}
	return tokenDuration, tokenString
}

// sendVerificationEmail sends an email with the verification link
func sendVerificationEmail(email, token string) {
	// TODO: Implement email sending logic here
	log.Debugf("Sending verification email to %s with OTP: %s", email, token)
}

func (u *Impl) UpdateMe(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{"message": "User successfully authenticated!"})
}

func (u *Impl) DeleteUser(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func populateUserEntry(newUser *models.User, req *dto.RegisterUserRequest, hashedPassword []byte) string {
	newUser.ID = uuid.New().String()
	newUser.Email = req.Email
	newUser.FirstName = req.FirstName
	newUser.LastName = req.LastName
	newUser.Password = string(hashedPassword)
	otp := common.GenerateOTP()
	newUser.VerificationCode = otp
	newUser.VerificationSentAt = time.Now()
	newUser.UpdatedAt = time.Now()
	newUser.CreatedAt = time.Now()
	return otp
}
