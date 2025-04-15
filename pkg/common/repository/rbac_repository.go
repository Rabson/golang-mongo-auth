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

var RbacRepo *BaseRepository[models.Rbac]

func SetRbacRepository(db *mongo.Database) {
	RbacRepo = NewRepository[models.Rbac](db, "casbin_rule")
}

func RbacFindById(id primitive.ObjectID, projection bson.M) (*models.Rbac, error) {

	if projection == nil {
		projection = bson.M{"_id": 1, "name": 1, "email": 1, "role": 1}
	}

	rbac, err := RbacRepo.FindOne(context.TODO(), bson.M{"_id": id}, projection)
	if err != nil {
		return nil, err
	}
	return rbac, nil
}

func RbacCreate(rbac models.Rbac) error {
	_, err := RbacRepo.InsertOne(context.TODO(), &rbac)
	if err != nil {
		return err
	}
	return nil
}
func RbacUpdateById(id primitive.ObjectID, rbac models.Rbac) error {
	var updateData = bson.M{}

	if rbac.PType != "" {
		updateData["ptype"] = rbac.PType
	}
	if rbac.V0 != "" {
		updateData["v0"] = rbac.V0
	}
	if rbac.V1 != "" {
		updateData["v1"] = rbac.V1
	}
	if rbac.V2 != "" {
		updateData["v2"] = rbac.V2
	}

	if len(updateData) == 0 {
		return errors.New("no data to update")
	}

	updateData["updatedAt"] = time.Now()

	update := bson.M{
		"$set": updateData,
	}
	_, err := RbacRepo.UpdateOne(context.TODO(), bson.M{"_id": id}, update)

	if err != nil {
		return err
	}
	return nil
}
