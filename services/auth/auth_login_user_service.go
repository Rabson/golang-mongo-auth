package auth_service

import (
	"encoding/hex"
	"errors"
	"log"
	"net/http"

	"golang-mongo-auth/models"
	"golang-mongo-auth/repository"
	"golang-mongo-auth/utils"
)

func LoginUser(data map[string]interface{}, userCtx models.UserCtx) (interface{}, error, int) {

	user := models.User{
		Email:    data["email"].(string),
		Password: data["password"].(string),
	}

	foundUser, foundUserErr := repository.UserFindByEmail(user.Email)

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
