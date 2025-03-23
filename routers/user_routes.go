package routers

import (
	"golang-mongo-auth/middleware"
	user_services "golang-mongo-auth/services/users"
	"golang-mongo-auth/validators"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine) {
	routeGroup := r.Group("/users")
	{
		routeGroup.GET("/details", middleware.AuthMiddleware(), middleware.ServiceWrapper(user_services.UserGetDetails, nil))
		routeGroup.GET("/:id", middleware.AuthMiddleware(), middleware.ServiceWrapper(user_services.UserGetDetails, nil))
		routeGroup.PUT("/", middleware.AuthMiddleware(), middleware.ServiceWrapper(user_services.UpdateUser, validators.UpdateProfileValidator{}))
	}
}
