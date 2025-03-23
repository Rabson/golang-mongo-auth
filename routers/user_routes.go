package routers

import (
	"golang-mongo-auth/middleware"
	"golang-mongo-auth/services"
	"golang-mongo-auth/validators"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine) {
	userGroup := r.Group("/users")
	{
		userGroup.GET("/:id?", middleware.AuthMiddleware(), middleware.ServiceWrapper(services.GetProfile, nil))
		userGroup.PUT("/", middleware.AuthMiddleware(), middleware.ServiceWrapper(services.UpdateProfile, validators.UpdateProfileValidator{}))
	}
}
