package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(rw http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(rw, req)
		return
	}
	rw.Write([]byte("<h1>Hello, welcome to GoSnippet</h1>"))
}

func showSnippet(rw http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.URL.Query().Get("id"))
	log.Println(id)
	if err != nil || id < 1 {
		http.NotFound(rw, req)
		return
	}
	//rw.Write([]byte("Display a specific snippet ..."))
	fmt.Fprintf(rw, "Display Snippet with specific ID %d", id)
}

func createSnippet(rw http.ResponseWriter, req *http.Request) {
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
		http.Error(rw, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	// setting and changing system generated headers
	rw.Header().Set("Content-Type", "application/json")

	//Suppressing system generated headers
	rw.Header()["Date"] = nil
	rw.Write([]byte("Create a new snippet..."))

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Starting server at port :4000")
	log.Fatal(http.ListenAndServe(":4000", mux))
}
