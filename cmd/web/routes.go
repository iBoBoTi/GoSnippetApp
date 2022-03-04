package main

import (
	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) Routes() http.Handler {

	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	dynamicMiddleware := alice.New(app.session.Enable)

	mux := pat.New()
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))
	mux.Get("/snippet/create", dynamicMiddleware.ThenFunc(app.createSnippetForm))
	mux.Post("/snippet/create", dynamicMiddleware.ThenFunc(app.createSnippet))
	mux.Get("/snippet/:id", dynamicMiddleware.ThenFunc(app.showSnippet))
	mux.Get("/user/signup", http.HandlerFunc(app.signupUserForm))
	mux.Post("/user/signup", http.HandlerFunc(app.signupUser))
	mux.Get("/user/login", http.HandlerFunc(app.loginUserForm))
	mux.Post("/user/login", http.HandlerFunc(app.loginUser))
	mux.Post("/user/logout", http.HandlerFunc(app.logoutUser))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))
	return standardMiddleware.Then(mux)
}
