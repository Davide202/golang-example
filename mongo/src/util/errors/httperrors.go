package errors

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"startwithmongo/util/logger"
)





func  ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	logger.Error().Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func  ClientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}