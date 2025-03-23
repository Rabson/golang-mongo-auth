package user_services

import (
	"golang-mongo-auth/models"
	"golang-mongo-auth/repository"
	"golang-mongo-auth/validators"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func UpdateUser(data map[string]interface{}, userCtx models.UserCtx) (interface{}, error, int) {

	var updateData validators.UpdateProfileValidator

	updateErr := repository.UserUpdateById(userCtx.UserId, bson.M{"name": updateData.Name})

	if updateErr != nil {
		log.Println("UpdateProfile: Error updating user:", updateErr.Error())
		return nil, updateErr, http.StatusInternalServerError
	}

	return nil, nil, 0
}
