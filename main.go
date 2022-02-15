package main

import (
	"log"
	"net/http"
)

func home(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("<h1>Hello, welcome to GoSnippet</h1>"))
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Println("Starting server at port :4000")
	log.Fatal(http.ListenAndServe(":4000", mux))
}
