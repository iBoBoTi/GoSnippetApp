package main

import (
	"bytes"
	"fmt"
	"github.com/justinas/nosurf"
	"net/http"
	"runtime/debug"
	"time"
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

func (app *application) render(rw http.ResponseWriter, req *http.Request, name string, td *templateData) {
	ts, ok := app.templateCache[name]
	if !ok {
		app.serverError(rw, fmt.Errorf("The template %s does not exist", name))
		return
	}

	buf := new(bytes.Buffer)

	err := ts.Execute(buf, app.addDefaultData(td, req))
	if err != nil {
		app.serverError(rw, err)
	}

	if _, err = buf.WriteTo(rw); err != nil {
		app.serverError(rw, err)
		return
	}
}

func (app *application) addDefaultData(td *templateData, req *http.Request) *templateData {
	if td == nil {
		td = &templateData{}
	}

	td.AuthenticatedUser = app.authenticatedUser(req)
	td.CSRFToken = nosurf.Token(req)
	td.Flash = app.session.PopString(req, "flash")
	td.CurrentYear = time.Now().Year()
	return td
}

func (app *application) authenticatedUser(r *http.Request) int {
	return app.session.GetInt(r, "userID")
}
