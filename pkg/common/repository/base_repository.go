package repository

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BaseRepository[T any] struct {
	Collection *mongo.Collection
}

func NewRepository[T any](db *mongo.Database, collectionName string) *BaseRepository[T] {
	return &BaseRepository[T]{
		Collection: db.Collection(collectionName),
	}
}

// InsertOne inserts a document into the collection
func (r *BaseRepository[T]) InsertOne(ctx context.Context, document *T) (*mongo.InsertOneResult, error) {
	result, err := r.Collection.InsertOne(ctx, document)
	if err != nil {
		log.Println("Error inserting document:", err)
	}
	return result, err
}

// FindOne finds a single document matching the filter
func (r *BaseRepository[T]) FindOne(ctx context.Context, filter bson.M, projection bson.M) (*T, error) {
	var result T

	opts := options.FindOne().SetProjection(projection)
	err := r.Collection.FindOne(ctx, filter, opts).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// FindAll finds all documents matching the filter
func (r *BaseRepository[T]) FindAll(ctx context.Context, filter bson.M) ([]T, error) {
	var results []T
	cursor, err := r.Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var elem T
		if err := cursor.Decode(&elem); err != nil {
			return nil, err
		}
		results = append(results, elem)
	}
	return results, nil
}

// UpdateOne updates a single document
func (r *BaseRepository[T]) UpdateOne(ctx context.Context, filter bson.M, update bson.M) (*mongo.UpdateResult, error) {
	result, err := r.Collection.UpdateOne(ctx, filter, bson.M{"$set": update})
	if err != nil {
		log.Println("Error updating document:", err)
	}
	return result, err
}

// DeleteOne deletes a single document
func (r *BaseRepository[T]) DeleteOne(ctx context.Context, filter bson.M) (*mongo.DeleteResult, error) {
	result, err := r.Collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Println("Error deleting document:", err)
	}
	return result, err
}
