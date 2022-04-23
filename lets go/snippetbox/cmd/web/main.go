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

	log.Println("Starting server on http://localhost:3000")
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}
