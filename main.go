package main

import (
	"golang-mongo-auth/database"
	"golang-mongo-auth/middleware"
	"golang-mongo-auth/routers"

	"github.com/gin-gonic/gin"
)

func main() {

	database.Init()

	r := gin.Default()

	routers.SetupRoutes(r)

	r.Use(middleware.ErrorHandler())

	r.Run(":8080")
}
