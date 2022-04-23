package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from snippetbox"))
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
