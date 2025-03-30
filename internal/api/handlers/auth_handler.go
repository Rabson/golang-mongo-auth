package handlers

import (
	"golang-mongo-auth/internal/api/middleware"
	"golang-mongo-auth/internal/api/request"
	"golang-mongo-auth/pkg/auth/service"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(r *gin.RouterGroup) {
	routeGroup := r.Group("/auth/v1")
	{
		routeGroup.POST("/register", middleware.ServiceWrapper(service.RegisterUser, request.RegisterValidator{}))
		routeGroup.POST("/login", middleware.ServiceWrapper(service.LoginUser, request.LoginValidator{}))
		routeGroup.POST("/admin-login", middleware.ServiceWrapper(service.LoginAdmin, request.LoginValidator{}))
	}
}
