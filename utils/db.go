package utils

import "go.mongodb.org/mongo-driver/mongo"

var UserColl *mongo.Collection

func SetUserCollection(collection *mongo.Collection) {
	UserColl = collection
}
