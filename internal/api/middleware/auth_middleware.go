package middleware

import (
	"golang-mongo-auth/pkg/common/constants"
	"golang-mongo-auth/pkg/common/types"
	"golang-mongo-auth/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(module types.Module, action types.Action) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		claims, ValidateTokenErr := utils.ValidateToken(tokenString)

		if ValidateTokenErr != nil {
			utils.ErrorResponse(c, http.StatusUnauthorized, ValidateTokenErr.Error())
			return
		}

		ObjectId, ObjectIdErr := utils.StringToObjectId(claims.UserId)

		if ObjectIdErr != nil {
			utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid user")
			return
		}

		allowedActions, exists := constants.RoleModuleActions[claims.Role][module]
		if !exists {
			utils.ErrorResponse(c, http.StatusForbidden, "Role not found")
			return
		}

		isAllowed := false
		for _, allowedAction := range allowedActions {
			if allowedAction == action {
				isAllowed = true
				break
			}
		}

		if !isAllowed {
			utils.ErrorResponse(c, http.StatusForbidden, "Access denied")
			return
		}

		c.Set("userCtx", types.UserCtx{
			UserId: ObjectId,
			Role:   claims.Role,
		})
		c.Next()
	}
}
