package utils

import "github.com/gin-gonic/gin"

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
