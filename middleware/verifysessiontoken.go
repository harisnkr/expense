package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	log "github.com/sirupsen/logrus"

	"github.com/harisnkr/expense/config"
)

// Claims structure to hold the information in the JWT token
type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

// Auth is a middleware to verify session tokens issued
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
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
			log.Fatal(err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token (ParseWithClaims)"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			c.Set("email", claims.Email)
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
	}
}
