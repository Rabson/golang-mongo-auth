package routers

import (
	"golang-mongo-auth/middleware"
	"golang-mongo-auth/services"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine) {
	r.GET("/profile", middleware.AuthMiddleware(), middleware.ServiceWrapper(services.GetProfile, nil))
}
