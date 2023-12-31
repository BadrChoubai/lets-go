package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"net/http"
	"snippetbox.badrchoubai.dev/ui"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})

	// Add a new GET /ping route.
	router.HandlerFunc(http.MethodGet, "/ping", ping)

	// Create file server to serve files out of "./ui/static"
	fileServer := http.FileServer(http.FS(ui.Files))
	// Create handler from the file server that serves all requests to /static/,
	// stripping "/static" before a request reaches the file server
	router.Handler(http.MethodGet, "/static/*filepath", fileServer)

	// Use the nosurf middleware on all our 'dynamic' routes
	dynamic := alice.New(app.sessionManager.LoadAndSave, app.authenticate, noSurf)

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

	// Because the 'protected' middlewhere chain appends to the 'dynamic'
	// the noSurf middleware will also be used on the three routes below
	protected := dynamic.Append(app.requireAuthentication)

	router.Handler(
		http.MethodGet,
		"/snippets/create",
		protected.ThenFunc(app.snippetCreate),
	)

	router.Handler(
		http.MethodPost,
		"/snippets/create",
		protected.ThenFunc(app.snippetCreatePost),
	)

	router.Handler(
		http.MethodPost,
		"/user/logout",
		protected.ThenFunc(app.userLogoutPost),
	)

	standard := alice.New(app.recoverAtPanic, app.logRequest, secureHeaders)

	return standard.Then(router)
}
