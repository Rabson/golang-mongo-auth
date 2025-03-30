package repository

import (
	"context"
	modelAdmin "golang-mongo-auth/pkg/admin/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var AdminRepo *BaseRepository[modelAdmin.Admin]

func SetAdminRepository(db *mongo.Database) {
	AdminRepo = NewRepository[modelAdmin.Admin](db, "Admins")
}

func AdminFindByEmail(email string) (*modelAdmin.Admin, error) {
	Admin, err := AdminRepo.FindOne(context.TODO(), bson.M{"email": email}, nil)
	if err != nil {
		return nil, err
	}
	return Admin, nil
}

func AdminFindById(id primitive.ObjectID, projection bson.M) (*modelAdmin.Admin, error) {

	if projection == nil {
		projection = bson.M{"_id": 1, "name": 1, "email": 1, "role": 1}
	}

	data, err := AdminRepo.FindOne(context.TODO(), bson.M{"id": id}, projection)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func AdminCreate(Admin modelAdmin.Admin) error {
	_, err := AdminRepo.InsertOne(context.TODO(), &Admin)
	if err != nil {
		return err
	}
	return nil
}
func AdminUpdateById(id primitive.ObjectID, Admin bson.M) error {
	_, err := AdminRepo.UpdateOne(context.TODO(), bson.M{"_id": id}, Admin)

	if err != nil {
		return err
	}
	return nil
}
