package repository

import (
	"context"
	"golang-mongo-auth/pkg/user/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var UserRepo *BaseRepository[models.User]

func SetUserRepository(db *mongo.Database) {
	UserRepo = NewRepository[models.User](db, "Users")
}

func UserFindByEmail(email string) (*models.User, error) {
	user, err := UserRepo.FindOne(context.TODO(), bson.M{"email": email}, nil)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UserFindById(id primitive.ObjectID, projection bson.M) (*models.User, error) {

	if projection == nil {
		projection = bson.M{"_id": 1, "name": 1, "email": 1, "role": 1}
	}

	user, err := UserRepo.FindOne(context.TODO(), bson.M{"_id": id}, projection)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UserCreate(user models.User) error {
	_, err := UserRepo.InsertOne(context.TODO(), &user)
	if err != nil {
		return err
	}
	return nil
}
func UserUpdateById(id primitive.ObjectID, user bson.M) error {
	_, err := UserRepo.UpdateOne(context.TODO(), bson.M{"_id": id}, user)

	if err != nil {
		return err
	}
	return nil
}
