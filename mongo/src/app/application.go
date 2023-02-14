package app

import (
	"log"
	"fmt"
	"net/http"
	"startwithmongo/repository"
	"startwithmongo/util/logger"
	"runtime/debug"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	bookings *repository.BookModel
}

func init(){
	var app application
	(&app).Set()
	
}

func (app *application) Set(){
	app.infoLog = logger.Info()
	app.errorLog = logger.Error()
}


func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}
