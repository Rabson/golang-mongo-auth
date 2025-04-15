package server

import (
	"golang-mongo-auth/internal/api/handlers"
	"golang-mongo-auth/internal/api/middleware"
	"golang-mongo-auth/pkg/common/database"
	"golang-mongo-auth/pkg/common/repository"
	"golang-mongo-auth/pkg/config"
	"golang-mongo-auth/pkg/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Start() {
	db := database.Init(config.GetMongoURI(), config.GetDbName())

	repository.SetRepositories(db)

	utils.LoadCasbinModel(config.GetMongoURI() + "/" + config.GetDbName())

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Content-Type", "Authorization"},
	}))

	handlers.SetupRoutes(r)

	r.Use(middleware.ErrorHandler())

	port := config.GetPort()
	if err := r.Run(":" + port); err != nil {
		panic(err)
	}
}
