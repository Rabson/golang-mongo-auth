package service

import (
	"golang-mongo-auth/internal/api/request"
	"golang-mongo-auth/pkg/common/repository"
	"golang-mongo-auth/pkg/common/types"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func UpdateUser(data map[string]interface{}, userCtx types.UserCtx) (interface{}, error, int) {

	var updateData request.UpdateProfileValidator

	// Validate and parse the input data
	if err := updateData.Validate(data); err != nil {
		log.Println("UpdateProfile: Validation error:", err.Error())
		return nil, err, http.StatusBadRequest
	}

	// Handle profile picture upload if provided
	if profile, ok := data["profile"].(string); ok && profile != "" {
		uploadedURL, uploadErr := uploadProfile(profile, userCtx.UserId.String())
		if uploadErr != nil {
			log.Println("UpdateProfile: Error uploading profile picture:", uploadErr.Error())
			return nil, uploadErr, http.StatusInternalServerError
		}
		updateData.Profile = uploadedURL
	}

	// Update the user data in the database
	updateErr := repository.UserUpdateById(userCtx.UserId, bson.M{
		"name":    updateData.Name,
		"profile": updateData.Profile,
	})

	if updateErr != nil {
		log.Println("UpdateProfile: Error updating user:", updateErr.Error())
		return nil, updateErr, http.StatusInternalServerError
	}

	return nil, nil, 0
}
