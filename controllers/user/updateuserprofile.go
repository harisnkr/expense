package user

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/harisnkr/expense/common"
	"github.com/harisnkr/expense/dto"
	"github.com/harisnkr/expense/models"
)

// UpdateProfile updates the authenticated user's profile
func (u *Impl) UpdateProfile(c *gin.Context) {
	var (
		log    = slog.With(common.RequestID, c.MustGet(common.RequestID))
		userID = c.Query("userID")
		req    *dto.UpdateMeRequest
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Debug("err", err, "invalid request body")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filter := bson.M{"userID": userID}
	update := bson.M{"$set": bson.M{}}

	if req.FirstName != nil {
		update["$set"].(bson.M)["first_name"] = *req.FirstName
	}
	if req.LastName != nil {
		update["$set"].(bson.M)["last_name"] = *req.LastName
	}
	if req.ProfilePicture != nil {
		update["$set"].(bson.M)["profile_picture"] = *req.ProfilePicture
	}

	result, err := u.collections.Users.UpdateOne(c, bson.M{"_id": userID}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	if result.ModifiedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found or no changes made"})
		return
	}

	var updatedUser models.User
	if err = u.collections.Users.FindOne(c, filter).Decode(&updatedUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch updated user"})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}
