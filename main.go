package main

import (
	"golang-mongo-auth/database"
	"golang-mongo-auth/middleware"
	"golang-mongo-auth/routers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	database.Init()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Content-Type", "Authorization"},
	}))

	routers.SetupRoutes(r)

	r.Use(middleware.ErrorHandler())

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
