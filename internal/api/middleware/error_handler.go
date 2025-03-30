package middleware

import (
	"golang-mongo-auth/pkg/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // Process request

		// Check for 404 errors
		if c.Writer.Status() == http.StatusNotFound {
			log.Printf("404 error: %d %s", c.Writer.Status(), c.Request.URL)
			utils.ErrorResponse(c, http.StatusNotFound, "Resource not found")
		}

		// Check for 500 errors
		if len(c.Errors) > 0 {
			log.Printf("500 error: %d %s", c.Writer.Status(), c.Request.URL)
			utils.ErrorResponse(c, http.StatusInternalServerError, "Internal server error")
		}
	}
}
