package app

import (
	"net/http"
	"startwithmongo/util/logger"
)

func loggingMiddleware1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		logger.Info().Println("FILTER 1")
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
func loggingMiddleware2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		logger.Info().Println("FILTER 2")
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
