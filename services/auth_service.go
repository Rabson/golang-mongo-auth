package services

import (
	"context"
	"encoding/hex"
	"errors"
	"log"
	"net/http"

	"golang-mongo-auth/models"
	"golang-mongo-auth/repository"
	"golang-mongo-auth/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegisterUser(data map[string]interface{}, userCtx models.UserCtx) (interface{}, error, int) {

	user := models.User{
		ID:       primitive.NewObjectID(),
		Name:     data["name"].(string),
		Email:    data["email"].(string),
		Password: data["password"].(string),
		Role:     models.RoleUser,
	}

	foundUser, foundUserErr := repository.UserRepo.FindOne(context.TODO(), bson.M{"email": user.Email}, nil)

	if foundUserErr != nil && foundUserErr.Error() != "mongo: no documents in result" {
		log.Println("RegisterUser: Error finding user:", foundUserErr.Error())
		return nil, errors.New("Something went wrong"), http.StatusInternalServerError
	}

	if foundUser != nil {
		return nil, errors.New("User already exists"), http.StatusConflict
	}

	hashedPassword, _ := utils.HashPassword(user.Password)

	user.Password = hashedPassword

	_, createUserErr := repository.UserRepo.InsertOne(context.TODO(), &user)

	if createUserErr != nil {
		log.Println("RegisterUser: Error creating user:", createUserErr.Error())
		return nil, errors.New("Failed to create user"), http.StatusInternalServerError

	}

	return map[string]interface{}{"message": "User created successfully"}, nil, 0
}

func LoginUser(data map[string]interface{}, userCtx models.UserCtx) (interface{}, error, int) {

	user := models.User{
		Email:    data["email"].(string),
		Password: data["password"].(string),
	}

	foundUser, foundUserErr := repository.UserRepo.FindOne(context.TODO(), bson.M{"email": user.Email}, nil)

	if foundUserErr != nil && foundUserErr.Error() != "mongo: no documents in result" {
		log.Println("RegisterUser: Error finding user:", foundUserErr.Error())
		return nil, errors.New("Something went wrong"), http.StatusInternalServerError
	}

	if foundUser == nil {
		log.Println("LoginUser: Error finding user:", foundUserErr)
		return nil, errors.New("Email not found"), http.StatusNotFound
	}

	if !utils.CheckPassword(foundUser.Password, user.Password) {
		return nil, errors.New("Invalid credentials"), http.StatusForbidden
	}

	var id string = hex.EncodeToString(foundUser.ID[:])

	token, _ := utils.GenerateToken(id)

	return map[string]interface{}{"token": token}, nil, 0
}
