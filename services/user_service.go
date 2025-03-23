package services

import (
	"context"
	"golang-mongo-auth/models"
	"golang-mongo-auth/repository"
	"golang-mongo-auth/validators"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func GetProfile(data map[string]interface{}, userCtx models.UserCtx) (interface{}, error, int) {

	projection := bson.M{"_id": 1, "name": 1, "email": 1, "role": 1}
	foundUser, foundUserErr := repository.UserRepo.FindOne(context.TODO(), bson.M{"_id": userCtx.UserId}, projection)

	if foundUserErr != nil {
		log.Println("GetProfile: Error finding user:", foundUserErr.Error())
		return nil, foundUserErr, 0
	}
	return foundUser, nil, 0
}

func UpdateProfile(data map[string]interface{}, userCtx models.UserCtx) (interface{}, error, int) {

	var updateData validators.UpdateProfileValidator

	_, updateErr := repository.UserRepo.UpdateOne(context.TODO(), bson.M{"_id": userCtx.UserId}, bson.M{
		"name": updateData.Name,
	})

	if updateErr != nil {
		log.Println("UpdateProfile: Error updating user:", updateErr.Error())
		return nil, updateErr, http.StatusInternalServerError
	}

	return nil, nil, 0
}
