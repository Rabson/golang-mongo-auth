package user_services

import (
	"golang-mongo-auth/models"
	"golang-mongo-auth/repository"
	"log"
)

func UserGetDetails(data map[string]interface{}, userCtx models.UserCtx) (interface{}, error, int) {

	foundUser, foundUserErr := repository.UserFindById(userCtx.UserId, nil)

	if foundUserErr != nil {
		log.Println("UserGetDetails: Error finding user:", foundUserErr.Error())
		return nil, foundUserErr, 0
	}
	return foundUser, nil, 0
}
