package services

import (
	"context"
	"encoding/hex"
	"errors"
	"log"

	"golang-mongo-auth/models"
	"golang-mongo-auth/repository"
	"golang-mongo-auth/utils"

	"go.mongodb.org/mongo-driver/bson"
)

func RegisterUser(data map[string]interface{}, userCtx models.UserCtx) (interface{}, error, *int) {

	user := models.User{
		Email:    data["email"].(string),
		Password: data["password"].(string),
	}

	foundUser, foundUserErr := repository.UserRepo.FindOne(context.TODO(), bson.M{"email": user.Email})

	if foundUserErr != nil && foundUserErr.Error() != "mongo: no documents in result" {
		log.Println("RegisterUser: Error finding user:", foundUserErr.Error())
		return nil, errors.New("Something went wrong"), nil
	}

	if foundUser != nil {
		return nil, errors.New("User already exists"), nil
	}

	hashedPassword, _ := utils.HashPassword(user.Password)
	user.Password = hashedPassword

	_, createUserErr := repository.UserRepo.InsertOne(context.TODO(), &user)

	if createUserErr != nil {
		return nil, errors.New("Failed to create user"), nil

	}

	return map[string]interface{}{"message": "User created successfully"}, nil, nil
}

func LoginUser(data map[string]interface{}, userCtx models.UserCtx) (interface{}, error, *int) {

	user := models.User{
		Email:    data["email"].(string),
		Password: data["password"].(string),
	}

	foundUser, foundUserErr := repository.UserRepo.FindOne(context.TODO(), bson.M{"email": user.Email})

	if foundUserErr != nil && foundUserErr.Error() != "mongo: no documents in result" {
		log.Println("RegisterUser: Error finding user:", foundUserErr.Error())
		return nil, errors.New("Something went wrong"), nil
	}

	if foundUser == nil {
		log.Println("LoginUser: Error finding user:", foundUserErr)
		return nil, errors.New("Email not found"), nil
	}

	if !utils.CheckPassword(foundUser.Password, user.Password) {
		return nil, errors.New("Invalid credentials"), nil
	}

	var id string = hex.EncodeToString(foundUser.ID[:])

	token, _ := utils.GenerateToken(id)

	return map[string]interface{}{"token": token}, nil, nil
}
