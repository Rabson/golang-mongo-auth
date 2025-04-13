package service

import (
	"errors"
	"golang-mongo-auth/pkg/common/constants"
	"golang-mongo-auth/pkg/common/messages"
	"golang-mongo-auth/pkg/common/repository"
	"golang-mongo-auth/pkg/common/types"

	"golang-mongo-auth/pkg/user/models"

	"golang-mongo-auth/pkg/utils"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegisterUser(data map[string]interface{}, userCtx types.UserCtx) (interface{}, error, int) {

	user := models.User{
		ID:       primitive.NewObjectID(),
		Name:     data["name"].(string),
		Email:    data["email"].(string),
		Password: data["password"].(string),
		Role:     constants.ROLE_USER,
	}

	foundUser, foundUserErr := repository.UserFindByEmail(user.Email)

	if foundUserErr != nil && foundUserErr.Error() != "mongo: no documents in result" {
		log.Println("RegisterUser: Error finding user:", foundUserErr.Error())
		return nil, errors.New(messages.ErrSomethingWentWrong), http.StatusInternalServerError
	}

	if foundUser != nil {
		return nil, errors.New(messages.ErrEmailAlreadyExists), http.StatusConflict
	}

	hashedPassword, _ := utils.HashPassword(user.Password)

	user.Password = hashedPassword

	createUserErr := repository.UserCreate(user)

	if createUserErr != nil {
		log.Println("RegisterUser: Error creating user:", createUserErr.Error())
		return nil, errors.New(messages.ErrOperationFailed), http.StatusInternalServerError

	}

	return nil, nil, http.StatusCreated
}
