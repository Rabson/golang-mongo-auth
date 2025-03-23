package routers

import (
	"golang-mongo-auth/middleware"
	auth_service "golang-mongo-auth/services/auth"
	"golang-mongo-auth/validators"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(r *gin.Engine) {

	routeGroup := r.Group("/auth")
	{
		routeGroup.POST("/register", middleware.ServiceWrapper(auth_service.RegisterUser, validators.RegisterValidator{}))
		routeGroup.POST("/login", middleware.ServiceWrapper(auth_service.LoginUser, validators.LoginValidator{}))
	}
}
