package main

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

type Route struct {
	Name       string `json:"name"`
	method     string
	Path       string `json:"path"`
	middleware alice.Chain
	handler    http.HandlerFunc
}

func newRoute(name, method, path string, middleware alice.Chain, handler http.HandlerFunc) *Route {
	return &Route{
		name,
		method,
		path,
		middleware,
		handler,
	}
}

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

	routes := []*Route{
		newRoute("index", http.MethodGet, "/", standardMiddleware, app.homeHandler),
		newRoute("links.new", http.MethodGet, "/link/new", standardMiddleware.Append(app.requireAuthentication), app.createHandler),
		newRoute("links.store", http.MethodPost, "/link/new", standardMiddleware.Append(app.requireAuthentication), app.storeHandler),
		newRoute("links.show", http.MethodGet, "/links/:id", standardMiddleware, app.showHandler),
		newRoute("links.tag", http.MethodGet, "/tags/:tag", standardMiddleware, app.byTagHandler),
		newRoute("auth.login", http.MethodGet, "/login", standardMiddleware, app.loginFormHandler),
		newRoute("auth.loginPost", http.MethodPost, "/login", standardMiddleware, app.loginHandler),
		newRoute("auth.logout", http.MethodPost, "/logout", standardMiddleware.Append(app.requireAuthentication), app.logoutHandler),
	}

	for _, r := range routes {
		router.Handler(r.method, r.Path, r.middleware.ThenFunc(r.handler))
	}

	fileServer := http.FileServer(http.Dir("ui/static/"))
	router.Handler(http.MethodGet, "/static/*path", standardMiddleware.Then(http.StripPrefix("/static", fileServer)))

	r, err := json.Marshal(routes)
	if err != nil {
		app.logger.Fatal("routes cannot be shared")
	}
	app.inertia.Share("routes", string(r))

	return router
}
