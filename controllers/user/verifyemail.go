package user

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/harisnkr/expense/dto"
	"github.com/harisnkr/expense/models"
)

// VerifyEmail verifies the email with the verification token
func (u *Impl) VerifyEmail(c *gin.Context) {
	var (
		collection = u.collections.Users
		log        = slog.With(c)
	)

	var req *dto.UserEmailVerifyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Fetch the user by email and verificationCode
	var user models.User
	err := collection.FindOne(c, bson.M{"email": req.Email, "verification_code": req.VerificationCode}).Decode(&user)
	if err != nil {
		log.Warn("verification code not found in database for given email")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid verification code"})
		return
	}

	// Mark the user as verified
	update := bson.M{"$set": bson.M{"verified": true}}
	if _, err = collection.UpdateOne(c, bson.M{"email": user.Email}, update); err != nil {
		log.Warn("failed to mark the user as verified: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Unknown error"})
		return
	}

	tokenDuration, tokenString := generateSessionJWT(c, user)
	if tokenString == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to generate session token"})
		return
	}
	c.JSON(http.StatusOK, UserLoginResponse{
		SessionToken: tokenString,
		ExpiresIn:    tokenDuration.String(),
	})
}
