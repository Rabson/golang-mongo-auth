package auth_service

import (
	"errors"
	"golang-mongo-auth/models"
	"golang-mongo-auth/repository"
	"golang-mongo-auth/utils"
	"log"
	"net/http"

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

	foundUser, foundUserErr := repository.UserFindByEmail(user.Email)

	if foundUserErr != nil && foundUserErr.Error() != "mongo: no documents in result" {
		log.Println("RegisterUser: Error finding user:", foundUserErr.Error())
		return nil, errors.New("Something went wrong"), http.StatusInternalServerError
	}

	if foundUser != nil {
		return nil, errors.New("User already exists"), http.StatusConflict
	}

	hashedPassword, _ := utils.HashPassword(user.Password)

	user.Password = hashedPassword

	createUserErr := repository.UserCreate(user)

	if createUserErr != nil {
		log.Println("RegisterUser: Error creating user:", createUserErr.Error())
		return nil, errors.New("Failed to create user"), http.StatusInternalServerError

	}

	return nil, nil, http.StatusCreated
}
