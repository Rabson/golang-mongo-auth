package service

import (
	"errors"
	"golang-mongo-auth/pkg/common/repository"
	"golang-mongo-auth/pkg/common/types"
	"golang-mongo-auth/pkg/fileManager"
	"log"
	"mime/multipart"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func UpdateUser(data map[string]interface{}, userCtx types.UserCtx) (interface{}, error, int) {

	type UpdateData struct {
		Name    string
		Profile string
	}

	type Data struct {
		data map[string]interface{}
	}

	var updateData UpdateData

	if name, ok := data["name"].(string); ok && name != "" {
		updateData.Name = name
	}

	if profile, ok := data["profile"].(*multipart.FileHeader); ok && profile != nil {
		uploadedURL, uploadErr := fileManager.UploadFile(
			fileManager.FILE_TYPE_PROFILE,
			profile,
			userCtx.UserId.Hex()+"_"+profile.Filename,
		)

		if uploadErr != nil {
			log.Println("UpdateProfile: Error uploading profile picture:", uploadErr.Error())
			return nil, errors.New("error uploading profile picture"), http.StatusInternalServerError
		}
		updateData.Profile = uploadedURL
	}

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
