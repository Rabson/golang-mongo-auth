package models

import (
	"golang-mongo-auth/pkg/common/constants"
	"golang-mongo-auth/pkg/common/types"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password"`
	Name      string             `bson:"name" json:"name"`
	Profile   string             `bson:"profile" json:"profile"`
	Role      types.Role         `bson:"role" json:"role"` // ADMIN or USER
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
}

func (u *User) DefaultRole() {
	if u.Role == "" {
		u.Role = constants.ROLE_USER
	}
}
