package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/footer.partial.tmpl",
		"./ui/html/base.layout.tmpl",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
	}

}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with id %d...", id)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		// w.WriteHeader(405)
		// w.Write([]byte("Method not allowed"))
		app.clientError(w, http.StatusBadGateway)
		return
	}
	w.Write([]byte("Create a new snippet..."))
}

func (app *application) editSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// w.WriteHeader(405)
		w.Header().Set("Allow", http.MethodPost) // Giving out which method is allowed
		w.Header().Set("Cache-Control", "public, max-age=31536000")
		w.Header().Get("Cache-Control")
		app.clientError(w, http.StatusPaymentRequired)
		return
	}
	w.Header().Get("Cache-Control")
	w.Header().Set("Content-Type", "application/json") // set the header Content-Type to json
	w.Write([]byte("{ 'snippet': 'bruh'}"))
	// w.Write([]byte("Editing snippets...."))
}
