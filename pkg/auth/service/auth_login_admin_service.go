package service

import (
	"golang-mongo-auth/pkg/common/constants"
	"golang-mongo-auth/pkg/common/types"
	"golang-mongo-auth/pkg/user/models"
)

func LoginAdmin(data map[string]interface{}, userCtx types.UserCtx) (interface{}, error, int) {

	user := models.User{
		Email:    data["email"].(string),
		Password: data["password"].(string),
	}

	userMap := map[string]interface{}{
		"email":     user.Email,
		"password":  user.Password,
		"loginRole": constants.ROLE_ADMIN,
	}

	return Login(userMap, userCtx)
}
