package database

import (
	"context"
	"golang-mongo-auth/pkg/common/repository"
	"golang-mongo-auth/pkg/config"

	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
)

func Init() {
	// Load environment variables from .env file
	if devEnvErr := godotenv.Load(); devEnvErr != nil {
		log.Println("Warning: No .env file found")
	}

	var mongoErr error
	client, mongoErr = mongo.Connect(context.TODO(), options.Client().ApplyURI(config.GetMongoURI()))
	if mongoErr != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", mongoErr)
	}
	db := client.Database(os.Getenv("DB_NAME"))
	repository.SetRepositories(db)
}
