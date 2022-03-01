package main

import (
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) Routes() http.Handler {

	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	return standardMiddleware.Then(mux)
}
