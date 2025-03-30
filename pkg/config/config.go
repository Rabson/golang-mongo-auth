package config

import (
	"os"
)

// GetMongoURI retrieves the MongoDB URI from environment variables
func GetMongoURI() string {
	return os.Getenv("MONGO_URI")
}
