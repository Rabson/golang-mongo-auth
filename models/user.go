package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role string

const (
	RoleAdmin Role = "ADMIN"
	RoleUser  Role = "USER"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"password"`
	Name     string             `bson:"name" json:"name"`
	Role     Role               `bson:"role" json:"role"` // ADMIN or USER
}

func (u *User) DefaultRole() {
	if u.Role == "" {
		u.Role = RoleUser
	}
}

func (u *User) ValidateRole() bool {
	return u.Role == RoleAdmin || u.Role == RoleUser
}
