package main

import (
	"golang-mongo-auth/internal/api/handlers"
	"golang-mongo-auth/internal/api/middleware"
	"golang-mongo-auth/pkg/common/database"
	"golang-mongo-auth/pkg/common/repository"
	"golang-mongo-auth/pkg/config"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if loadEnvErr := godotenv.Load(); loadEnvErr != nil {
		log.Println("Warning: No .env file found")
	}

	db := database.Init(config.GetMongoURI(), config.GetDbName())

	repository.SetRepositories(db)

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
