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
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}

func (app *application) home(rw http.ResponseWriter, req *http.Request) {

	if req.URL.Path != "/" {
		//http.NotFound(rw, req)
		app.notFound(rw)
		return
	}

	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(rw, err)
		return
	}

	data := &templateData{Snippets: s}

	files := []string{
		"./ui/html/home.page.gohtml",
		"./ui/html/base.layout.gohtml",
		"./ui/html/footer.partial.gohtml",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		//http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		app.serverError(rw, err)
		return
	}
	err = ts.Execute(rw, data)
	if err != nil {
		app.errorLog.Println(err.Error())
		//http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		app.serverError(rw, err)
	}
}

func (app *application) showSnippet(rw http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.URL.Query().Get("id"))
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

	data := &templateData{
		Snippet: s,
	}

	files := []string{
		"./ui/html/show.page.gohtml",
		"./ui/html/base.layout.gohtml",
		"./ui/html/footer.partial.gohtml",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		//http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		app.serverError(rw, err)
		return
	}
	err = ts.Execute(rw, data)
	if err != nil {
		app.errorLog.Println(err.Error())
		//http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		app.serverError(rw, err)
	}
}

func (app *application) createSnippet(rw http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		// tells what methods are allowed
		rw.Header().Set("Allow", "POST")
		rw.Header().Add("Content-Type", "application/json")

		//// writes the header
		//rw.WriteHeader(http.StatusMethodNotAllowed)
		//
		//// writes to the response body
		//rw.Write([]byte("Method Not Allowed\n"))

		// shortcut for the above code --> combines WriteHeader method and Write method for non 200 status code
		//http.Error(rw, "Method Not Allowed", http.StatusMethodNotAllowed)
		app.clientError(rw, http.StatusMethodNotAllowed)
		return
	}
	// setting and changing system generated headers
	rw.Header().Set("Content-Type", "application/json") // default text/plain

	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi"
	expires := "7"

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(rw, err)
	}
	http.Redirect(rw, req, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)

	//Suppressing system generated headers
	//rw.Header()["Date"] = nil
	//rw.Write([]byte("Create a new snippet..."))

}
