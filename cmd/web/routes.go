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

	dynamic := alice.New(app.sessionManager.LoadAndSave)

	router.Handler(
		http.MethodGet,
		"/",
		dynamic.ThenFunc(app.home),
	)

	router.Handler(
		http.MethodGet,
		"/snippets/view/:id",
		dynamic.ThenFunc(app.snippetView),
	)

	router.Handler(
		http.MethodGet,
		"/snippets/create",
		dynamic.ThenFunc(app.snippetCreate),
	)

	router.Handler(
		http.MethodPost,
		"/snippets/create",
		dynamic.ThenFunc(app.snippetCreatePost),
	)

	router.HandlerFunc(
		http.MethodDelete,
		"/snippets/:id",
		app.snippetDelete,
	)

	router.Handler(
		http.MethodGet,
		"/user/signup",
		dynamic.ThenFunc(app.userSignup),
	)

	router.Handler(
		http.MethodPost,
		"/user/signup",
		dynamic.ThenFunc(app.userSignupPost),
	)

	router.Handler(
		http.MethodGet,
		"/user/login",
		dynamic.ThenFunc(app.userLogin),
	)

	router.Handler(
		http.MethodPost,
		"/user/login",
		dynamic.ThenFunc(app.userLoginPost),
	)

	router.Handler(
		http.MethodPost,
		"/user/logout",
		dynamic.ThenFunc(app.userLogoutPost),
	)

	standard := alice.New(app.recoverAtPanic, app.logRequest, secureHeaders)

	return standard.Then(router)
}
