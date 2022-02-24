package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *application) serverError(rw http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)
	http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(rw http.ResponseWriter, status int) {
	http.Error(rw, http.StatusText(status), status)
}

func (app *application) notFound(rw http.ResponseWriter) {
	app.clientError(rw, http.StatusNotFound)
}
