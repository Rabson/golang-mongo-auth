package handlers

import (
	"context"
	"net/http"

	"golang-mongo-auth/models"
	"golang-mongo-auth/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetProfile(c *gin.Context) {
	userId, _ := c.Get("userId")
	var user models.User

	err := utils.UserColl.FindOne(context.TODO(), bson.M{"_id": userId}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}
