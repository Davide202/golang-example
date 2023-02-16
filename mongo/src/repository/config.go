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

func (db *DatabaseConfiguration) Set(
		Database *mongo.Database, ctx context.Context,
		) {
	db.Database = Database
	db.ctx = ctx
}

func getCollection(collection string) (*mongo.Collection) {
	var coll *mongo.Collection
	_ , ok := DB.ctx.Deadline()
	if ok {
		coll = DB.Database.Collection(collection)
	}else{
		//todo 
	}
	return coll
}