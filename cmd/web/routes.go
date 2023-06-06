package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	http "net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})

	// Create file server to serve files out of "./ui/static"
	fileServer := http.FileServer(http.Dir("./ui/static"))
	// Create handler from the file server that serves all requests to /static/,
	// stripping "/static" before a request reaches the file server
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	router.HandlerFunc(
		http.MethodGet,
		"/",
		app.home,
	)

	router.HandlerFunc(
		http.MethodGet,
		"/snippets/view/:id",
		app.snippetView,
	)

	router.HandlerFunc(
		http.MethodGet,
		"/snippets/create",
		app.snippetCreate,
	)

	router.HandlerFunc(
		http.MethodPost,
		"/snippets/create",
		app.snippetCreatePost,
	)

	router.HandlerFunc(
		http.MethodDelete,
		"/snippets/:id",
		app.snippetDelete,
	)

	standard := alice.New(app.recoverAtPanic, app.logRequest, secureHeaders)

	return standard.Then(router)
}
