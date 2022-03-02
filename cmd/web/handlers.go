package main

import (
	"fmt"
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
	snippets      *mysql.SnippetModel
	templateCache map[string]*template.Template
}

type templateData struct {
	CurrentYear int
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
	//rw.Write([]byte("Display a specific snippet ..."))
	//fmt.Fprintf(rw, "%v", s)
	app.render(rw, req, "show.page.gohtml", &templateData{Snippet: s})
}

func (app *application) createSnippet(rw http.ResponseWriter, req *http.Request) {
	// setting and changing system generated headers
	err := req.ParseForm()
	if err != nil {
		app.clientError(rw, http.StatusBadRequest)
	}

	title := req.PostForm.Get("title")
	content := req.PostForm.Get("content")
	expires := req.PostForm.Get("expires")

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(rw, err)
	}
	http.Redirect(rw, req, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)

	//Suppressing system generated headers
	//rw.Header()["Date"] = nil
	//rw.Write([]byte("Create a new snippet..."))

}

func (app *application) createSnippetForm(rw http.ResponseWriter, req *http.Request) {
	app.render(rw, req, "create.page.gohtml", nil)
}
