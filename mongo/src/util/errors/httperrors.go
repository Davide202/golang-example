package errors

import (
	"net/http"
	"startwithmongo/util/logger"
)

func ServerError(w http.ResponseWriter, err error) {
	//trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	//logger.Error().Output(2, trace)
	logger.Error().Println(err.Error())
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func NotFoundError(w http.ResponseWriter, err error) {
	logger.Error().Println(err.Error())
	http.Error(w, "Not Found", http.StatusNotFound)
}

func BadRequestError(w http.ResponseWriter, err error) {
	logger.Error().Println(err.Error())
	http.Error(w, err.Error(), http.StatusBadRequest)
}

func ClientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func HttpError(w http.ResponseWriter, status int, err error) {
	logger.Error().Println(err.Error())
	http.Error(w, http.StatusText(status), status)
}
