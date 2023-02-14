package repository

import (
	"context"
	"errors"
	"startwithmongo/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


type BookModel struct {
	C *mongo.Collection
}

func (m *BookModel) findAll() ([]model.Book,error){
	ctx := context.TODO()
	b := []model.Book{}

	cursor,err := m.C.Find(ctx,bson.M{})
	if err != nil {return nil,err}
	
	//prima opzione
	for cursor.Next(ctx){
		var book model.Book
		cursor.Decode(&book)//should hanlde error 
		b = append(b, book)
	}
	//seconda opzione
	cursor.All(ctx,&b)//should hanlde error 

	return b,nil
}

func (m *BookModel) FindById(id string)(*model.Book,error){
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var book = model.Book{}
	err = m.C.FindOne(context.TODO(), bson.M{"_id": p}).Decode(&book)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}

	return &book, nil
}


func (m *BookModel) Insert(booking model.Book) (*mongo.InsertOneResult, error) {
	return m.C.InsertOne(context.TODO(), booking)
}


func (m *BookModel) Delete(id string) (*mongo.DeleteResult, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return m.C.DeleteOne(context.TODO(), bson.M{"_id": p})
}