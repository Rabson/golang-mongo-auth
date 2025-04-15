package middleware

import (
	"golang-mongo-auth/pkg/common/constants"
	"golang-mongo-auth/pkg/common/messages"
	"golang-mongo-auth/pkg/common/types"
	"golang-mongo-auth/pkg/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(module types.Module) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		method := types.HttpMethod(c.Request.Method)

		claims, ValidateTokenErr := utils.ValidateToken(tokenString)
		if ValidateTokenErr != nil {
			utils.ErrorResponse(c, http.StatusUnauthorized, ValidateTokenErr.Error())
			return
		}

		log.Println("AuthMiddleware: claims userId", claims.UserId)
		log.Println("AuthMiddleware: claims Role", claims.Role)
		log.Println("AuthMiddleware: module", module)
		log.Println("AuthMiddleware: method", method)

		ObjectId, ObjectIdErr := utils.StringToObjectId(claims.UserId)
		if ObjectIdErr != nil {
			utils.ErrorResponse(c, http.StatusUnauthorized, messages.ErrAccessDenied)
			return
		}

		action, exists := constants.HttpMethodToAction[method]
		if !exists {
			utils.ErrorResponse(c, http.StatusBadRequest, messages.ErrPermissionDenied)
			return
		}

		// allowedActions, exists := constants.roleModuleActions[claims.Role][module]
		// if !exists {
		// 	utils.ErrorResponse(c, http.StatusForbidden, messages.ErrInvalidRole)
		// 	return
		// }

		// isAllowed := false
		// for _, allowedAction := range allowedActions {
		// 	if allowedAction == action {
		// 		isAllowed = true
		// 		break
		// 	}
		// }

		isAllowed := utils.ValidateCasbin(string(claims.Role), module, action)

		if !isAllowed {
			utils.ErrorResponse(c, http.StatusForbidden, messages.ErrAccessDenied)
			return
		}

		c.Set("userCtx", types.UserCtx{
			UserId: ObjectId,
			Role:   claims.Role,
		})
		c.Next()
	}
}
