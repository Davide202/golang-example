package repository

import (
	"errors"
	"startwithmongo/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const collection = "books"

func FindAll() (*[]model.Book, error) {

	b := []model.Book{}

	cursor, err := getCollection(collection).Find(DB.ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	//prima opzione
	for cursor.Next(DB.ctx) {
		var book model.Book
		cursor.Decode(&book) //should hanlde error
		b = append(b, book)
	}
	//seconda opzione
	cursor.All(DB.ctx, &b) //should hanlde error

	return &b, nil
}

func FindById(p primitive.ObjectID) (*model.Book, error) {

	var book = model.Book{}
	coll := getCollection(collection)

	err := coll.FindOne(DB.ctx, bson.M{"_id": p}).Decode(&book)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}

	return &book, nil
}

func FindByTitle(title string) (*[]model.Book, error) {
	var books = make([]model.Book, 0)

	filter := bson.D{{
		Key: "title", Value: primitive.Regex{
			Pattern: "[a-zA-Z0-9]+[" + title + "][a-zA-Z0-9]+",
			Options: "i",
		}}}
	/*filter := bson.D{{"title", bson.D{{"$all", bson.A{title}}}}}

	filter := bson.D{{"title", bson.D{{"$text", bson.A{title}}}}}


	filter := bson.D{{
		Key: "title",
		Value: bson.D{{"$regex", bson.A{title}}}}}
	*/

	coll := getCollection(collection)

	cursor, err := coll.Find(DB.ctx, filter)
	if err != nil {
		return nil, err
	}
	//prima opzione
	for cursor.Next(DB.ctx) {
		var book model.Book
		cursor.Decode(&book) //should hanlde error
		books = append(books, book)
	}
	return &books, nil
}

func Insert(booking model.Book) (*mongo.InsertOneResult, error) {

	return getCollection(collection).InsertOne(DB.ctx, booking)
}

func Delete(id string) (*mongo.DeleteResult, error) {

	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return getCollection(collection).DeleteOne(DB.ctx, bson.M{"_id": p})
}
