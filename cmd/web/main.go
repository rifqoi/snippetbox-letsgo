package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)
	mux.HandleFunc("/snippet/edit", editSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// StripPrefix buat nge strip /static
	// e.g /static/image.jpg bakal distrip jadi -> /image.jpg
	// dari situ ntar server bakal ngirim file dari ".ui/static/image.jpg"
	// https://stackoverflow.com/questions/27945310/why-do-i-need-to-use-http-stripprefix-to-access-my-static-files
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Starting server on http://localhost:3000")
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}
