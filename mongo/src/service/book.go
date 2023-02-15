package service

import(
	"startwithmongo/model"
	"startwithmongo/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindAllBooks()([]*model.BookDTO,error){
	_ , err := repository.FindAll()
	
	return nil,err
}

func FindBookByID(id string	)(*model.BookDTO,error){

	return nil,nil
}

func InsertOneBook(book model.BookDTO) (*mongo.InsertOneResult,error){

	return nil, nil
}

func DeleteBookById(id string )(*mongo.DeleteResult, error){

	return nil,nil
}