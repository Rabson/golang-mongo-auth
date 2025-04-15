package service

import (
	"errors"
	"golang-mongo-auth/pkg/common/repository"
	"golang-mongo-auth/pkg/common/types"
	"golang-mongo-auth/pkg/fileManager"
	"golang-mongo-auth/pkg/user/models"
	"log"
	"mime/multipart"
	"net/http"
)

func UpdateUser(data map[string]interface{}, userCtx types.UserCtx) (interface{}, error, int) {

	type UpdateData struct {
		Name    string
		Profile string
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

	updateErr := repository.UserUpdateById(userCtx.UserId, models.User{
		Name:    updateData.Name,
		Profile: updateData.Profile,
	})

	if updateErr != nil {
		log.Println("UpdateProfile: Error updating user:", updateErr.Error())
		return nil, updateErr, http.StatusInternalServerError
	}

	return nil, nil, 0
}
