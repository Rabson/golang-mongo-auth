package service

import (
	"errors"
	"golang-mongo-auth/pkg/common/messages"
	"golang-mongo-auth/pkg/common/repository"
	"golang-mongo-auth/pkg/common/types"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UserGetDetails(data map[string]interface{}, userCtx types.UserCtx) (interface{}, error, int) {

	var userId primitive.ObjectID
	id, ok := data["id"].(string)
	if ok {
		var objErr error
		userId, objErr = primitive.ObjectIDFromHex(id)
		if objErr != nil {
			return nil, errors.New(messages.ErrSomethingWentWrong), 0
		}
	} else {
		userId = userCtx.UserId
	}

	objectID := userId

	foundUser, foundUserErr := repository.UserFindById(objectID, nil)

	if foundUserErr != nil {
		log.Println("UserGetDetails: Error finding user:", foundUserErr.Error())
		return nil, errors.New(messages.ErrUserNotFound), 404
	}
	return foundUser, nil, 0
}
