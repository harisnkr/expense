package user

import (
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

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
	Login(ctx *gin.Context)
	UpdateProfile(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
	GetEmailOTP(ctx *gin.Context)
}

// Impl is the implementation for user.API
type Impl struct {
	database    *mongo.Client
	collections *data.Collections
}

// New creates and returns a new user.API implementation for usage with routes
func New(database *mongo.Client, collections *data.Collections) *Impl {
	return &Impl{database, collections}
}

// DeleteUser deletes the authenticated user's profile
func (u *Impl) DeleteUser(_ *gin.Context) {
	//TODO implement me
	panic("implement me")
}

// GetEmailOTP is an internal endpoint (used for testing) to retrieve OTP assigned to user's email
func (u *Impl) GetEmailOTP(c *gin.Context) {
	var (
		email = c.Query("email")
		log   = slog.With(common.RequestID, c.MustGet(common.RequestID))
	)
	emailEscaped, _ := url.QueryUnescape(email)
	log = log.With("email", emailEscaped)
	log.Info("getting email OTP")

	var user models.User
	err := u.collections.Users.FindOne(c, bson.M{"email": strings.TrimSpace(emailEscaped)}).Decode(&user)
	if err != nil {
		log.Warn("OTP not found for given email")
		c.JSON(http.StatusNotFound, gin.H{"message": "Unknown error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"otp": user.VerificationCode})
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
	newUser.Cards = []models.Card{}
	newUser.Transactions = []models.Transaction{}
	newUser.Budgets = []models.Budget{}
	newUser.Savings = []models.Savings{}

	return otp
}

func generateSessionJWT(c *gin.Context, user models.User) (time.Duration, string) {
	var (
		log      = slog.With(common.RequestID, c.MustGet(common.RequestID)).With("userID", user.ID, "email", user.Email)
		tokenTTL = config.SessionTokenTTLInHours
	)

	exp := time.Now().Add(config.SessionTokenTTLInHours).Unix()
	iat := time.Now().Unix()
	nbf := time.Now().Unix()
	iss := common.Issuer
	sub := user.ID
	aud := user.ID

	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"email": user.Email,
		"jti":   uuid.New().String(),
		"exp":   exp,
		"iat":   iat,
		"iss":   iss,
		"nbf":   nbf,
		"sub":   sub,
		"aud":   aud,
	})
	tokenString, err := token.SignedString(config.ECDSAKey)
	if err != nil {
		log.Error("failed to generate jwt", err)
		return time.Duration(0), tokenString
	}
	return tokenTTL, tokenString
}

func sendVerificationEmail(c *gin.Context, email, token string) {
	// TODO: Implement email sending logic here]
	slog.InfoContext(c, fmt.Sprintf("Sending verification email to %s with OTP: %s", email, token))
}
