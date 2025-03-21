package handlers

import (
	"context"
	"net/http"

	"golang-mongo-auth/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func RegisterUser(c *gin.Context) {
	var user utils.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	hashedPassword, _ := utils.HashPassword(user.Password)
	user.Password = hashedPassword

	_, err := utils.UserColl.InsertOne(context.TODO(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func LoginUser(c *gin.Context) {
	var creds utils.User
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var user utils.User
	err := utils.UserColl.FindOne(context.TODO(), bson.M{"email": creds.Email}).Decode(&user)
	if err != nil || !utils.CheckPassword(user.Password, creds.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, _ := utils.GenerateToken(user.Email)
	c.JSON(http.StatusOK, gin.H{"token": token})
}
