package utils

import (
	"errors"
	"golang-mongo-auth/pkg/common/messages"
	"golang-mongo-auth/pkg/common/types"
	"log"

	"github.com/gin-gonic/gin"
)

func ErrorResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{
		"error":   true,
		"message": message,
	})
	c.Abort()
}

func SuccessResponse(c *gin.Context, data interface{}, statusCode int) {
	if data == nil {
		data = gin.H{}
	}
	c.JSON(statusCode, gin.H{
		"error":   false,
		"message": "Success",
		"data":    data,
	})
	c.Abort()
}

func GetUserContext(c *gin.Context) (types.UserCtx, error) {
	userContext, exists := c.Get("userCtx")
	var user types.UserCtx
	if exists {
		var ok bool
		user, ok = userContext.(types.UserCtx)
		if !ok {
			log.Println("Failed to cast userCtx to UserCtx")
			return types.UserCtx{}, errors.New(messages.ErrSomethingWentWrong)
		}
	}
	return user, nil
}
