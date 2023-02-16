package controller

import (
	"encoding/json"
	"net/http"
	"startwithmongo/model"
	"github.com/gorilla/mux"
	"startwithmongo/util/errors"
	"startwithmongo/util/logger"
	"startwithmongo/service"
	
)



func  All(w http.ResponseWriter, r *http.Request) {
	// Get all bookings stored
	/**/ bookings, err := service.FindAllBooks()
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

func  FindByID(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	// Find booking by id
	m, err := service.FindBookByID(id)
	if err != nil {
		if err.Error() == "ErrNoDocuments" {
			logger.Info().Println("Book not found")
			return
		}
		// Any other error will send an internal server error
		errors.ServerError(w, err)
	}

	// Convert booking to json encoding
	b, err := json.Marshal(m)
	if err != nil {
		errors.ServerError(w, err)
	}

	logger.Info().Println("Have been found a booking")

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
	w.WriteHeader(http.StatusNoContent)
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