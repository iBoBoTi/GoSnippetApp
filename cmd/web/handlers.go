package main

import (
	"fmt"
	"github.com/golangcollege/sessions"
	"github.com/iBoBoTi/go-snippet/pkg/forms"
	"github.com/iBoBoTi/go-snippet/pkg/models"
	"github.com/iBoBoTi/go-snippet/pkg/models/mysql"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	session       *sessions.Session
	snippets      *mysql.SnippetModel
	templateCache map[string]*template.Template
	users         *mysql.UserModel
}

type templateData struct {
	CurrentYear int
	Flash       string
	Form        *forms.Form
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
}

func (app *application) home(rw http.ResponseWriter, req *http.Request) {
	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(rw, err)
		return
	}

	app.render(rw, req, "home.page.gohtml", &templateData{
		Snippets: s,
	})
}

func (app *application) showSnippet(rw http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		//http.NotFound(rw, req)
		app.notFound(rw)
		return
	}
	s, err := app.snippets.Get(id)
	if err == models.ErrNoRecord {
		app.notFound(rw)
		return
	} else if err != nil {
		app.serverError(rw, err)
		return
	}
	app.render(rw, req, "show.page.gohtml", &templateData{
		Snippet: s,
	})
}

func (app *application) createSnippet(rw http.ResponseWriter, req *http.Request) {
	// setting and changing system generated headers
	err := req.ParseForm()
	if err != nil {
		app.clientError(rw, http.StatusBadRequest)
	}

	form := forms.New(req.PostForm)
	form.Required("title", "content", "expires")
	form.MaxLength("title", 100)
	form.PermittedValues("expires", "365", "7", "1")

	if !form.Valid() {
		app.render(rw, req, "create.page.gohtml", &templateData{Form: form})
		return
	}

	id, err := app.snippets.Insert(form.Get("title"), form.Get("content"), form.Get("expires"))
	if err != nil {
		app.serverError(rw, err)
	}

	app.session.Put(req, "flash", "Snippet successfully created!")
	http.Redirect(rw, req, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}

func (app *application) createSnippetForm(rw http.ResponseWriter, req *http.Request) {
	app.render(rw, req, "create.page.gohtml", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) signupUserForm(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Display the user signup form...")
}
func (app *application) signupUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create a new user...")
}
func (app *application) loginUserForm(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Display the user login form...")
}
func (app *application) loginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Authenticate and login the user...")
}
func (app *application) logoutUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Logout the user...")
}
