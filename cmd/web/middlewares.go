package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/matteoggl/fireside/internal/data"
)

func (app *application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.logger.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())
		next.ServeHTTP(w, r)
	})
}

func secureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("X-Frame-Options", "deny")

		next.ServeHTTP(w, r)
	})
}

func (app *application) requireAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !app.isAuthenticated(r) {
			app.session.Put(r, "redirectPathAfterLogin", r.URL.Path)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		w.Header().Add("Cache-Control", "no-store")

		next.ServeHTTP(w, r)
	})
}

func (app *application) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		exists := app.session.Exists(r, "authenticatedUserID")

		if !exists {
			app.inertia.Share(string(ctxKeyIsAuthenticated), false)
			next.ServeHTTP(w, r)
			return
		}

		_, err := app.models.Users.Get(app.session.GetInt(r, "authenticatedUserID"))
		if errors.Is(err, data.ErrNoRecord) {
			app.session.Remove(r, "authenticatedUserID")
			app.inertia.Share(string(ctxKeyIsAuthenticated), false)
			next.ServeHTTP(w, r)
			return
		} else if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}

		ctx := context.WithValue(r.Context(), ctxKeyIsAuthenticated, true)
		app.inertia.Share(string(ctxKeyIsAuthenticated), true)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				app.serverErrorResponse(w, r, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}
