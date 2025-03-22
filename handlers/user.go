package handlers

import (
	"context"
	"net/http"

	"golang-mongo-auth/repository"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetProfile(c *gin.Context) {
	userId, _ := c.Get("userId")

	foundUser, foundUserErr := repository.UserRepo.FindOne(context.TODO(), bson.M{"_id": userId})

	if foundUserErr != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, foundUser)
}
