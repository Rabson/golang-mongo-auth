package handlers

import (
	"golang-mongo-auth/internal/api/middleware"
	"golang-mongo-auth/internal/api/request"
	"golang-mongo-auth/pkg/common/constants"
	"golang-mongo-auth/pkg/user/service"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.RouterGroup) {
	routeGroup := r.Group("/users/v1")
	{
		routeGroup.GET("/details", middleware.AuthMiddleware(constants.MODULE_USER), middleware.ServiceWrapper(service.UserGetDetails, nil))
		routeGroup.GET("/:id", middleware.AuthMiddleware(constants.MODULE_USER), middleware.ServiceWrapper(service.UserGetDetails, nil))
		routeGroup.PUT("/", middleware.AuthMiddleware(constants.MODULE_USER), middleware.ServiceWrapper(service.UpdateUser, request.UpdateProfileValidator{}))
	}
}
