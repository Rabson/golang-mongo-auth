package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserCtx struct {
	UserId primitive.ObjectID
	Role   Role
}

type Role string

type Module string

type Action string

type RoleModuleAction map[Role]map[Module][]Action
