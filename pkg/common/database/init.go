package database

import (
	"context"

	"log"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
)

func Init(url string, dbName string) *mongo.Database {
	// Load environment variables from .env file
	if devEnvErr := godotenv.Load(); devEnvErr != nil {
		log.Println("Warning: No .env file found")
	}

	var mongoErr error
	client, mongoErr = mongo.Connect(context.TODO(), options.Client().ApplyURI(url))
	if mongoErr != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", mongoErr)
	}
	db := client.Database(dbName)

	return db
}
