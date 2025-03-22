package repository

import (
	"golang-mongo-auth/models"

	"go.mongodb.org/mongo-driver/mongo"
)

var UserRepo *MongoRepository[models.User]

func SetUserRepository(db *mongo.Database) {
	UserRepo = NewMongoRepository[models.User](db, "users")
}
