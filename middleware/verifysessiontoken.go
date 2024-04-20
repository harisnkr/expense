package middleware

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"github.com/harisnkr/expense/common"
	"github.com/harisnkr/expense/config"
)

// Claims structure to hold the information in the JWT token
type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

// Auth is a middleware to verify session tokens issued
func Auth() gin.HandlerFunc {
	if common.GetMode() == common.Development {
		// skip auth middleware if development environment
		return func(c *gin.Context) {
			slog.Warn("skipping auth middleware in development env")
			c.Set(common.UserID, "")
			c.Next()
		}
	}
	return func(c *gin.Context) {
		log := slog.With(common.RequestID, c.MustGet(common.RequestID))
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return &config.ECDSAKey.PublicKey, nil
		})
		if err != nil {
			log.Error("jwt.ParseWithClaims failed", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token (ParseWithClaims)"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			c.Set("email", claims.Email)
			c.Set("userID", claims.Subject)
			log.With(common.Email, claims.Email, common.UserID, claims.Subject).
				Info("User authenticated")
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
	}
}
