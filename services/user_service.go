package services

import (
	"context"
	"golang-mongo-auth/models"
	"golang-mongo-auth/repository"

	"go.mongodb.org/mongo-driver/bson"
)

func GetProfile(data map[string]interface{}, userCtx models.UserCtx) (interface{}, error, *int) {

	foundUser, foundUserErr := repository.UserRepo.FindOne(context.TODO(), bson.M{"_id": userCtx.UserId})

	if foundUserErr != nil {
		return nil, foundUserErr, nil
	}
	return foundUser, nil, nil
}
