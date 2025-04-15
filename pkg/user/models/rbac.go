package models

import (
	"golang-mongo-auth/pkg/common/types"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Rbac struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	PType     string             `bson:"ptype" json:"ptype"`
	V0        types.Role         `bson:"v0" json:"v0"`
	V1        types.Module       `bson:"v1" json:"v1"`
	V2        types.Action       `bson:"v2" json:"v2"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
}
