package middleware

import (
	"golang-mongo-auth/internal/interfaces"
	"golang-mongo-auth/pkg/common/messages"
	"golang-mongo-auth/pkg/common/types"
	"golang-mongo-auth/pkg/utils"
	"log"
	"net/http"

	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func ServiceWrapper(serviceFunc func(map[string]interface{}, types.UserCtx) (interface{}, error, int), validator interfaces.Validator) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := utils.GetUserContext(c)
		if err != nil {
			utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		requestData := make(map[string]any)
		if c.Request.Method == http.MethodGet {
			if err := c.Request.ParseForm(); err != nil {
				utils.ErrorResponse(c, http.StatusBadRequest, messages.ErrInvalidData)
				return
			}
			for key, values := range c.Request.URL.Query() {
				if len(values) > 0 {
					requestData[key] = values[0]
				}
			}
		} else if c.Request.Method == http.MethodPost || c.Request.Method == http.MethodPut || c.Request.Method == http.MethodPatch {
			contentType := c.GetHeader("Content-Type")
			if contentType == "application/json" {
				if err := c.ShouldBindJSON(&requestData); err != nil {
					utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
					return
				}
			} else if contentType == "application/x-www-form-urlencoded" || strings.Contains(contentType, "multipart/form-data") {
				if err := c.ShouldBindWith(&requestData, binding.FormMultipart); err != nil {
					utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
					return
				}

				keys := validator.GetKeys()

				if len(keys) > 0 {
					for _, key := range keys {
						if value, ok := c.GetPostForm(key); ok {
							requestData[key] = value
						}
					}

					for _, key := range keys {
						if value, err := c.FormFile(key); err == nil {
							log.Println("File found:", value.Filename)
							requestData[key] = value
						}
					}
				}

			} else {
				utils.ErrorResponse(c, http.StatusUnsupportedMediaType, "Unsupported Content-Type")
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
