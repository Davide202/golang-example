package service

import (
	"startwithmongo/model"
	"startwithmongo/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindAllBooks() (*[]model.BookDTO, error) {
	all, err := repository.FindAll()
	if err != nil {
		return nil, err
	}
	dto := model.ToDto(all)
	if err != nil {
		return nil, err
	}
	return dto, nil
}

func FindBookByID(id string) (*model.BookDTO, error) {

	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	book, err := repository.FindById(p)
	if err != nil {
		return nil, err
	}
	dto := book.ToDto()
	return dto, nil
}

func FindByTitle(title string)(*[]model.BookDTO,error){
	all, err := repository.FindByTitle(title)
	if err != nil {
		return nil, err
	}
	dto := model.ToDto(all)
	if err != nil {
		return nil, err
	}
	return dto, nil
}

func InsertOneBook(dto model.BookDTO) (*mongo.InsertOneResult, error) {
	var err error
	book := dto.ToEntity()
	/* NON SERVE L'ID SI AUTOGENERA
	id := random.RandomHexStringFromUUID()
	book.ID, err = primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.Info().Println("ERROR creating book id " + err.Error())
	} */
	result, err := repository.Insert(book)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func RandomHexStringFromUUID() {
	panic("unimplemented")
}

func DeleteBookById(id string) (*mongo.DeleteResult, error) {
	result, err := repository.Delete(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}
