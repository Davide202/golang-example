package app

import (
	"startwithmongo/controller"
	"github.com/gorilla/mux"
)


func Routes() *mux.Router {
	// Register handler functions.
	r := mux.NewRouter()
	r.HandleFunc("/api/bookings/", controller.All).Methods("GET")
	r.HandleFunc("/api/bookings/{id}", controller.FindByID).Methods("GET")
	r.HandleFunc("/api/bookings/title/{title}", controller.FindByTitle).Methods("GET")
	r.HandleFunc("/api/bookings/", controller.Insert).Methods("POST")
	r.HandleFunc("/api/bookings/{id}", controller.Delete).Methods("DELETE")

	return r
}
