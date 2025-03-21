package main

import (
	"context"
	"log"
	"os"

	"golang-mongo-auth/handlers"
	"golang-mongo-auth/middleware"
	"golang-mongo-auth/utils"

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
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found")
	}

	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal(err)
	}
	userColl = client.Database("testdb").Collection("users")
	utils.SetUserCollection(userColl)

	// Set secret key from environment variables
	// utils.SetJWTKey(os.Getenv("JWT_SECRET"))
}

func main() {
	r := gin.Default()
	r.POST("/register", handlers.RegisterUser)
	r.POST("/login", handlers.LoginUser)
	r.GET("/profile", middleware.AuthMiddleware(), handlers.GetProfile)

	r.Run(":8080")
}
