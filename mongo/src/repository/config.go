package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type DatabaseConfiguration struct {
	Database *mongo.Database
	ctx      context.Context
}

var DB DatabaseConfiguration

func (db *DatabaseConfiguration) Set(Database *mongo.Database, ctx context.Context) {
	db.Database = Database
	db.ctx = ctx
}
