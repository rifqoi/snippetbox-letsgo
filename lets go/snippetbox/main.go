package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/" {
    http.NotFound(w, r)
    return
  }
  w.Write([]byte("Hello from snippetbox"))
}

func showSnippet(w http.ResponseWriter, r *http.Request){
  w.Write([]byte("Display a specific snippets"))
}

func createSnippet(w http.ResponseWriter, r *http.Request){
  w.Write([]byte("Create a new snippet..."))
}

func main(){
  mux := http.NewServeMux()
  mux.HandleFunc("/", home)
  mux.HandleFunc("/snippet", showSnippet)
  mux.HandleFunc("/snippet/create", createSnippet)

  log.Println("Starting server on :3000")
  err := http.ListenAndServe(":3000", mux)
  log.Fatal(err)
}
