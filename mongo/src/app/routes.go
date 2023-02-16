package app

import (
	"net/http"
	"startwithmongo/controller"
	"startwithmongo/util/logger"

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
	r.Use(loggingMiddleware)
	return r
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		logger.Info().Println(r.RequestURI)
		bearerToken := r.Header.Get("Authorization")
		if verifyToken(bearerToken) {
			// Call the next handler, which can be another middleware in the chain, or the final handler.
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}

	})
}

func verifyToken(bearerToken string) bool {
	return true
}
