package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Repository defines the standard CRUD operations
type Repository[T any] interface {
	InsertOne(ctx context.Context, document *T) (*mongo.InsertOneResult, error)
	FindOne(ctx context.Context, filter bson.M) (*T, error)
	FindAll(ctx context.Context, filter bson.M) ([]T, error)
	UpdateOne(ctx context.Context, filter bson.M, update bson.M) (*mongo.UpdateResult, error)
	DeleteOne(ctx context.Context, filter bson.M) (*mongo.DeleteResult, error)
}
