package routers

import (
	"golang-mongo-auth/middleware"
	"golang-mongo-auth/services"
	"golang-mongo-auth/validators"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(r *gin.Engine) {
	r.POST("/register", middleware.ServiceWrapper(services.RegisterUser, validators.RegisterValidator{}))
	r.POST("/login", middleware.ServiceWrapper(services.LoginUser, validators.LoginValidator{}))
}
