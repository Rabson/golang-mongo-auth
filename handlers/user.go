package handlers

import (
	"context"
	"net/http"

	"golang-mongo-auth/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetProfile(c *gin.Context) {
	email, _ := c.Get("email")
	var user utils.User

	err := utils.UserColl.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}
