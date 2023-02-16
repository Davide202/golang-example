package app

import (
	"net/http"
	"startwithmongo/controller"

	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	// Register handler functions.
	r := mux.NewRouter()
	r.HandleFunc("/api/bookings/", controller.All).Methods(http.MethodGet)
	r.HandleFunc("/api/bookings/{id}", controller.FindByID).Methods(http.MethodGet)
	r.HandleFunc("/api/bookings/title/{title}", controller.FindByTitle).Methods(http.MethodGet)
	r.HandleFunc("/api/bookings/", controller.Insert).Methods(http.MethodPost)
	r.HandleFunc("/api/bookings/{id}", controller.Delete).Methods(http.MethodDelete)

	r.HandleFunc("/health", controller.HealthCheckHandler)

	r.Use(loggingMiddleware1,loggingMiddleware2)
	//https://pkg.go.dev/github.com/gorilla/mux@v1.8.0#CORSMethodMiddleware
	r.Use(mux.CORSMethodMiddleware(r))
	return r
}
