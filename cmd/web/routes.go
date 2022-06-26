package main

import "net/http"

func (app *application) routes() *http.ServeMux {

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)
	mux.HandleFunc("/snippet/edit", app.editSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// StripPrefix buat nge strip /static
	// e.g /static/image.jpg bakal distrip jadi -> /image.jpg
	// dari situ ntar server bakal ngirim file dari ".ui/static/image.jpg"
	// https://stackoverflow.com/questions/27945310/why-do-i-need-to-use-http-stripprefix-to-access-my-static-files
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
