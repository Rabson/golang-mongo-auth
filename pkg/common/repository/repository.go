package repository

import "go.mongodb.org/mongo-driver/mongo"

func SetRepositories(db *mongo.Database) {
	SetUserRepository(db)
	SetAdminRepository(db)
}
