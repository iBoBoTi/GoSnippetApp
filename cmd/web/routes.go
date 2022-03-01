package main

import "net/http"

func (app *application) Routes() http.Handler {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	return app.logRequest(secureHeaders(mux))
}
