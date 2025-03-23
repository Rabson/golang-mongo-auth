package repository

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoRepository is a generic MongoDB repository
type MongoRepository[T any] struct {
	collection *mongo.Collection
}

// NewMongoRepository creates a new repository instance
func NewMongoRepository[T any](db *mongo.Database, collectionName string) *MongoRepository[T] {
	return &MongoRepository[T]{
		collection: db.Collection(collectionName),
	}
}

// InsertOne inserts a document into the collection
func (r *MongoRepository[T]) InsertOne(ctx context.Context, document *T) (*mongo.InsertOneResult, error) {
	result, err := r.collection.InsertOne(ctx, document)
	if err != nil {
		log.Println("Error inserting document:", err)
	}
	return result, err
}

// FindOne finds a single document matching the filter
func (r *MongoRepository[T]) FindOne(ctx context.Context, filter bson.M, projection bson.M) (*T, error) {
	var result T

	opts := options.FindOne().SetProjection(projection)
	err := r.collection.FindOne(ctx, filter, opts).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// FindAll finds all documents matching the filter
func (r *MongoRepository[T]) FindAll(ctx context.Context, filter bson.M) ([]T, error) {
	var results []T
	cursor, err := r.collection.Find(ctx, filter)
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
func (r *MongoRepository[T]) UpdateOne(ctx context.Context, filter bson.M, update bson.M) (*mongo.UpdateResult, error) {
	result, err := r.collection.UpdateOne(ctx, filter, bson.M{"$set": update})
	if err != nil {
		log.Println("Error updating document:", err)
	}
	return result, err
}

// DeleteOne deletes a single document
func (r *MongoRepository[T]) DeleteOne(ctx context.Context, filter bson.M) (*mongo.DeleteResult, error) {
	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Println("Error deleting document:", err)
	}
	return result, err
}
