package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserCtx struct {
	UserId primitive.ObjectID
}
