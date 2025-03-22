package utils

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var UserColl *mongo.Collection

func SetUserCollection(collection *mongo.Collection) {
	UserColl = collection
}

func StringToObjectId(id string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(id)
}
