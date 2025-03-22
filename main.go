package main

import (
	"context"
	"log"
	"os"

	"golang-mongo-auth/handlers"
	"golang-mongo-auth/middleware"
	"golang-mongo-auth/repository"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client   *mongo.Client
	userColl *mongo.Collection
)

func init() {
	// Load environment variables from .env file
	if devEnvErr := godotenv.Load(); devEnvErr != nil {
		log.Println("Warning: No .env file found")
	}

	var mongoErr error
	client, mongoErr = mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if mongoErr != nil {
		log.Fatal(mongoErr)
	}
	database := client.Database(os.Getenv("DB_NAME"))
	repository.SetUserRepository(database)
}

func main() {
	r := gin.Default()
	r.POST("/register", handlers.RegisterUser)
	r.POST("/login", handlers.LoginUser)
	r.GET("/profile", middleware.AuthMiddleware(), handlers.GetProfile)

	r.Run(":8080")
}
