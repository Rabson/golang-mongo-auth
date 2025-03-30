package models

import (
	"golang-mongo-auth/pkg/common/types"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Admin struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"password"`
	Name     string             `bson:"name" json:"name"`
	Role     types.Role         `bson:"role" json:"role"` // ADMIN or USER
}
