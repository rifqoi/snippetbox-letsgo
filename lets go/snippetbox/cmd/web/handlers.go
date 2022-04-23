package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	ts, err := template.ParseFiles("./ui/html/home.page.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 502)
	}

	err = ts.Execute(w, nil)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 502)

	}

}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with id %d...", id)
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		// w.WriteHeader(405)
		// w.Write([]byte("Method not allowed"))
		http.Error(w, "Method not allowed", 405)
		return
	}
	w.Write([]byte("Create a new snippet..."))
}

func editSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// w.WriteHeader(405)
		w.Header().Set("Allow", http.MethodPost) // Giving out which method is allowed
		w.Header().Set("Cache-Control", "public, max-age=31536000")
		w.Header().Get("Cache-Control")
		http.Error(w, "Try again bruh", 405)
		return
	}
	w.Header().Get("Cache-Control")
	w.Header().Set("Content-Type", "application/json") // set the header Content-Type to json
	w.Write([]byte("{ 'snippet': 'bruh'}"))
	// w.Write([]byte("Editing snippets...."))
}
