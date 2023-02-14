package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct{
	ID         	primitive.ObjectID `bson:"_id,omitempty"`
	Title     	string             `bson:"title"`
	Pages 		int             	`bson:"pages"`
	Kind     	[]string           `bson:"kind"`
}