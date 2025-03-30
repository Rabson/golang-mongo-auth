package service

import (
	"encoding/hex"
	"errors"

	modelAdmin "golang-mongo-auth/pkg/admin/models"

	modelUser "golang-mongo-auth/pkg/user/models"

	"golang-mongo-auth/pkg/common/constants"
	"golang-mongo-auth/pkg/common/repository"
	"golang-mongo-auth/pkg/common/types"
	"golang-mongo-auth/pkg/utils"

	"log"
	"net/http"
)

func Login(data map[string]interface{}, userCtx types.UserCtx) (interface{}, error, int) {

	loginRole, ok := data["loginRole"].(types.Role)
	if !ok {
		return nil, errors.New("Invalid login role"), http.StatusBadRequest
	}

	email, emailOk := data["email"].(string)
	password, passOk := data["password"].(string)

	if !emailOk || !passOk {
		return nil, errors.New("Invalid data"), http.StatusBadRequest
	}

	var foundUser *modelUser.User
	var foundAdmin *modelAdmin.Admin
	var foundErr error
	switch loginRole {
	case constants.ROLE_ADMIN:
		{
			foundAdmin, foundErr = repository.AdminFindByEmail(email)

		}
	case constants.ROLE_USER:
		{
			foundUser, foundErr = repository.UserFindByEmail(email)

		}
	}

	if foundErr != nil && foundErr.Error() != "mongo: no documents in result" {
		log.Println("RegisterUser: Error finding user:", foundErr.Error())
		return nil, errors.New("Something went wrong"), http.StatusInternalServerError
	}

	if foundUser == nil && foundAdmin == nil {
		return nil, errors.New("Email not found"), http.StatusNotFound
	}

	var checkPassword string

	if foundUser != nil {
		checkPassword = foundUser.Password
	} else if foundAdmin != nil {
		checkPassword = foundAdmin.Password
	}

	if !utils.CheckPassword(checkPassword, password) {
		return nil, errors.New("Invalid credentials"), http.StatusForbidden
	}

	var id string = hex.EncodeToString(foundUser.ID[:])

	token, _ := utils.GenerateToken(id, loginRole)

	return map[string]interface{}{"token": token}, nil, 0
}
