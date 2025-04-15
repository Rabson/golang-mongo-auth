package repository

import (
	"context"
	"errors"
	"golang-mongo-auth/pkg/user/models"
	"time"

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
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now
	_, err := UserRepo.InsertOne(context.TODO(), &user)
	if err != nil {
		return err
	}
	return nil
}
func UserUpdateById(id primitive.ObjectID, user models.User) error {
	var updateData = bson.M{}

	if user.Role != "" {
		updateData["role"] = user.Role
	}
	if user.Name != "" {
		updateData["name"] = user.Name
	}

	if user.Profile != "" {
		updateData["profile"] = user.Profile
	}
	if user.Email != "" {
		updateData["email"] = user.Email
	}
	if user.Password != "" {
		updateData["password"] = user.Password
	}
	if user.CreatedAt != (time.Time{}) {
		updateData["createdAt"] = user.CreatedAt
	}

	if len(updateData) == 0 {
		return errors.New("no data to update")
	}

	updateData["updatedAt"] = time.Now()

	update := bson.M{
		"$set": updateData,
	}
	_, err := UserRepo.UpdateOne(context.TODO(), bson.M{"_id": id}, update)

	if err != nil {
		return err
	}
	return nil
}
