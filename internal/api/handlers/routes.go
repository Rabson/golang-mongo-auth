package handlers

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	routeGroup := r.Group("/api")
	SetupAuthRoutes(routeGroup)
	SetupUserRoutes(routeGroup)
}
