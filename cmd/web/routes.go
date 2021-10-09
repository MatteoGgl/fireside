package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.logger.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())
		next.ServeHTTP(w, r)
	})
}

func (app *application) routes() *httprouter.Router {
	standardMiddleware := alice.New(app.logRequest, app.inertia.Middleware)

	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)

	router.Handler(http.MethodGet, "/", standardMiddleware.ThenFunc(app.homeHandler))
	router.Handler(http.MethodGet, "/link/:id", standardMiddleware.ThenFunc(app.showHandler))
	router.Handler(http.MethodGet, "/tag/:tag", standardMiddleware.ThenFunc(app.byTagHandler))

	router.Handler(http.MethodGet, "/login", standardMiddleware.ThenFunc(app.loginFormHandler))
	router.Handler(http.MethodPost, "/login", standardMiddleware.ThenFunc(app.loginHandler))

	fileServer := http.FileServer(http.Dir("ui/static/"))
	router.Handler(http.MethodGet, "/static/*path", standardMiddleware.Then(http.StripPrefix("/static", fileServer)))

	return router
}
