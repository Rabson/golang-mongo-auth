package routers

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	SetupAuthRoutes(r)
	SetupUserRoutes(r)
}
