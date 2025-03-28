package middleware

import (
	"golang-mongo-auth/interfaces"
	"golang-mongo-auth/models"
	"golang-mongo-auth/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ServiceWrapper(serviceFunc func(map[string]interface{}, models.UserCtx) (interface{}, error, int), validator interfaces.Validator) gin.HandlerFunc {
	return func(c *gin.Context) {
		userData, exists := c.Get("userCtx")
		var user models.UserCtx
		if exists {
			var ok bool
			user, ok = userData.(models.UserCtx)
			if !ok {
				log.Println("Failed to cast userCtx to UserCtx")
				utils.ErrorResponse(c, http.StatusInternalServerError, "Something went wrong")
				return
			}
		}

		requestData := make(map[string]any)
		if c.Request.Method == http.MethodGet {
			if err := c.Request.ParseForm(); err != nil {
				utils.ErrorResponse(c, http.StatusBadRequest, "Failed to parse form data")
				return
			}
			for key, values := range c.Request.URL.Query() {
				if len(values) > 0 {
					requestData[key] = values[0]
				}
			}
		} else if c.Request.Method == http.MethodPost || c.Request.Method == http.MethodPut || c.Request.Method == http.MethodPatch {
			if err := c.ShouldBindJSON(&requestData); err != nil {
				// if err := c.ShouldBindWith(&requestData, binding.FormMultipart); err != nil {
				utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
				// c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"
				return
			}
		}

		for _, param := range c.Params {
			requestData[param.Key] = param.Value
		}

		if validator != nil && requestData != nil {
			if err := validator.Validate(requestData); err != nil {
				utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
				return
			}
		}

		if result, err, statusCode := serviceFunc(requestData, user); err != nil {
			utils.ErrorResponse(c, statusCode, err.Error())
		} else {
			utils.SuccessResponse(c, result, func() int {
				if statusCode != 0 {
					return statusCode
				}
				return http.StatusOK
			}())
		}
	}
}
