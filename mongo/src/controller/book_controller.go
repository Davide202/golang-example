package controller

import (
	"encoding/json"
	"net/http"
	"startwithmongo/model"
	"startwithmongo/service"
	"startwithmongo/util/errors"
	"startwithmongo/util/logger"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func All(w http.ResponseWriter, r *http.Request) {
	// Get all bookings stored
	/**/
	bookings, err := service.FindAllBooks()
	if err != nil {
		errors.ServerError(w, err)
	}

	// Convert booking list into json encoding
	b, err := json.Marshal(bookings)
	if err != nil {
		errors.ServerError(w, err)
	}

	logger.Info().Println("Books have been listed")

	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func FindByID(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	// Find book by id
	m, err := service.FindBookByID(id)
	if err != nil {
		if err.Error() == "ErrNoDocuments" {
			logger.Info().Println("Book not found")
			errors.NotFoundError(w, err)
			return
		}
		if err == primitive.ErrInvalidHex {
			errors.BadRequestError(w, err)
			return
		}
		// Any other error will send an internal server error
		errors.ServerError(w, err)
		return
	}

	// Convert booking to json encoding
	b, err := json.Marshal(m)
	if err != nil {
		logger.Error().Println("Error Marshalling")
		errors.ServerError(w, err)
		return
	}

	logger.Info().Println("Have been found a book")

	// Send response back //
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func FindByTitle(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	title := vars["title"]
	logger.Info().Println("Looking for book by title: " + title)
	/**/
	bookings, err := service.FindByTitle(title)
	if err != nil {
		errors.ServerError(w, err)
		return
	}

	// Convert booking list into json encoding
	b, err := json.Marshal(bookings)
	if err != nil {
		errors.ServerError(w, err)
		return
	}

	logger.Info().Println("Books have been listed ")

	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	// Define booking model
	var m model.BookDTO
	// Get request information
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		errors.ServerError(w, err)
	}

	// Insert new booking
	insertResult, err := service.InsertOneBook(m)
	if err != nil {
		errors.ServerError(w, err)
	}

	logger.Info().Printf("New book have been created, id=%s", insertResult.InsertedID)
	w.WriteHeader(http.StatusCreated)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	// Delete booking by id
	deleteResult, err := service.DeleteBookById(id)
	if err != nil {
		errors.ServerError(w, err)
	}

	logger.Info().Printf("Have been eliminated %d booking(s)", deleteResult.DeletedCount)
	w.WriteHeader(http.StatusNoContent)
}
