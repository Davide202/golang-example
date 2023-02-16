package repository

import (
	
	"errors"
	"startwithmongo/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)
const collection = "books"

func FindAll() (*[]model.Book,error){

	b := []model.Book{}

	cursor,err := getCollection(collection).Find( DB.ctx,bson.M{})
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

	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var book = model.Book{}
	err =  getCollection(collection).FindOne( DB.ctx, bson.M{"_id": p}).Decode(&book)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}

	return &book, nil
}


func  Insert(booking model.Book) (*mongo.InsertOneResult, error) {
	
	return getCollection(collection).InsertOne(DB.ctx, booking)
}


func  Delete(id string) (*mongo.DeleteResult, error) {
	
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return getCollection(collection).DeleteOne(DB.ctx, bson.M{"_id": p})
}