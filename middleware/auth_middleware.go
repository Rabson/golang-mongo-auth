package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"golang-mongo-auth/models"
	"golang-mongo-auth/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}

		claims, err := utils.ParseToken(tokenString)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		ObjectId, ObjectIdErr := utils.StringToObjectId(claims.UserId)

		if ObjectIdErr != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user"})
			c.Abort()
			return
		}

		c.Set("userCtx", models.UserCtx{UserId: ObjectId})
		c.Next()
	}
}
