package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type Books struct{
	Books []Book
}

type Book struct{
	ID         	primitive.ObjectID `bson:"_id,omitempty"`
	Title     	string             `bson:"title"`
	Pages 		int             	`bson:"pages"`
	Kind     	[]string           `bson:"kind"`
}

type BookDTO struct{
	Title     	string             `json:"title"`
	Pages 		int             	`json:"pages"`
	Kind     	[]string           `json:"kind"`
}

func (book *Book) ToDto() (BookDTO){
	return BookDTO{
		Title: book.Title,
		Pages: book.Pages,
		Kind: book.Kind,
	}
}
func (books *Books) ToDto () ([]BookDTO){
	
}

func (dto *BookDTO) ToEntity() (Book){
	return Book{
		Title: dto.Title,
		Pages: dto.Pages,
		Kind: dto.Kind,
	}
}