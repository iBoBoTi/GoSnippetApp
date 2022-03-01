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

func (app *application) render(rw http.ResponseWriter, req http.Request, name string, td templateData) {
	ts, ok := app.templateCache[name]
	if !ok {
		app.serverError(rw, fmt.Errorf("The template %s oes not exist", name))
		return
	}
	err := ts.Execute(rw, td)
	if err != nil {
		app.serverError(rw, err)
	}
}
