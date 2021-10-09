package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)


func (app *application) routes() *httprouter.Router {
	standardMiddleware := alice.New(
		app.recoverPanic,
		app.logRequest,
		app.inertia.Middleware,
		secureHeaders,
		app.session.Enable,
		app.authenticate,
	)

	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)

	router.Handler(http.MethodGet, "/", standardMiddleware.ThenFunc(app.homeHandler))
	router.Handler(http.MethodGet, "/link/:id", standardMiddleware.ThenFunc(app.showHandler))
	router.Handler(http.MethodGet, "/tag/:tag", standardMiddleware.ThenFunc(app.byTagHandler))

	router.Handler(http.MethodGet, "/login", standardMiddleware.ThenFunc(app.loginFormHandler))
	router.Handler(http.MethodPost, "/login", standardMiddleware.ThenFunc(app.loginHandler))
	router.Handler(http.MethodPost, "/logout", standardMiddleware.Append(app.requireAuthentication).ThenFunc(app.logoutHandler))

	fileServer := http.FileServer(http.Dir("ui/static/"))
	router.Handler(http.MethodGet, "/static/*path", standardMiddleware.Then(http.StripPrefix("/static", fileServer)))

	return router
}
