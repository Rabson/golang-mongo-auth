package handlers

import (
	"context"
	"encoding/hex"
	"log"
	"net/http"

	"golang-mongo-auth/models"
	"golang-mongo-auth/repository"
	"golang-mongo-auth/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func RegisterUser(c *gin.Context) {
	var user models.User
	if jsonValidateErr := c.ShouldBindJSON(&user); jsonValidateErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	foundUser, foundUserErr := repository.UserRepo.FindOne(context.TODO(), bson.M{"email": user.Email})

	if foundUserErr != nil && foundUserErr.Error() != "mongo: no documents in result" {
		log.Println("RegisterUser: Error finding user:", foundUserErr.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}

	if foundUser != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}

	hashedPassword, _ := utils.HashPassword(user.Password)
	user.Password = hashedPassword

	_, createUserErr := repository.UserRepo.InsertOne(context.TODO(), &user)

	if createUserErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func LoginUser(c *gin.Context) {
	var creds models.User
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	foundUser, foundUserErr := repository.UserRepo.FindOne(context.TODO(), bson.M{"email": creds.Email})

	if foundUserErr != nil && foundUserErr.Error() != "mongo: no documents in result" {
		log.Println("RegisterUser: Error finding user:", foundUserErr.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}

	if foundUser == nil {
		log.Println("LoginUser: Error finding user:", foundUserErr)
		c.JSON(http.StatusForbidden, gin.H{"error": "Email not found"})
		return
	}

	if !utils.CheckPassword(foundUser.Password, creds.Password) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Invalid credentials"})
		return
	}

	var id string = hex.EncodeToString(foundUser.ID[:])

	token, _ := utils.GenerateToken(id)

	c.JSON(http.StatusOK, gin.H{"token": token})
}
