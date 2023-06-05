package main

import (
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	// Create file server to serve files out of "./ui/static"
	fileServer := http.FileServer(http.Dir("./ui/static"))
	// Create handler from the file server that serves all requests to /static/,
	// stripping "/static" before a request reaches the file server
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippets/view", app.snippetView)
	mux.HandleFunc("/snippets/create", app.snippetCreate)

	standard := alice.New(app.recoverAtPanic, app.logRequest, secureHeaders)

	return standard.Then(mux)
}
