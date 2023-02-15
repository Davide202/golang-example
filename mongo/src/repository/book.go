package repository

import (
	"context"
	"errors"
	"startwithmongo/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)
const collection = "books"
var BD DatabaseConfiguration

type BookModel struct {
	C *mongo.Collection
	Ctx context.Context
}

func FindAll() (*[]model.Book,error){

	var coll *mongo.Collection
	_ , ok := DB.ctx.Deadline()
	if ok {
		coll = DB.Database.Collection(collection)
	}else{
		//todo 
	}
	
	
	b := []model.Book{}

	cursor,err := coll.Find( DB.ctx,bson.M{})
	if err != nil {return nil,err}
	
	//prima opzione
	for cursor.Next( DB.ctx){
		var book model.Book
		cursor.Decode(&book)//should hanlde error 
		b = append(b, book)
	}
	//seconda opzione
	cursor.All( DB.ctx,&b)//should hanlde error 

	return &b,nil
}

func  FindById(id string)(*model.Book,error){

	var coll *mongo.Collection
	_ , ok := DB.ctx.Deadline()
	if ok {
		coll = DB.Database.Collection(collection)
	}else{
		//todo 
	}

	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var book = model.Book{}
	err =  coll.FindOne( DB.ctx, bson.M{"_id": p}).Decode(&book)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}

	return &book, nil
}


func  Insert(booking model.Book) (*mongo.InsertOneResult, error) {
	var coll *mongo.Collection
	_ , ok := DB.ctx.Deadline()
	if ok {
		coll = DB.Database.Collection(collection)
	}else{
		//todo 
	}
	return coll.InsertOne(DB.ctx, booking)
}


func  Delete(id string) (*mongo.DeleteResult, error) {
	var coll *mongo.Collection
	_ , ok := DB.ctx.Deadline()
	if ok {
		coll = DB.Database.Collection(collection)
	}else{
		//todo 
	}
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return coll.DeleteOne(DB.ctx, bson.M{"_id": p})
}