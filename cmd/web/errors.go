package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *application) logError(r *http.Request, err error) {
	app.logger.Println(err)
}

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	w.Header().Set("Status", fmt.Sprintf("%d", status))

	err := app.inertia.Render(w, r, "Error", map[string]interface{}{
		"error": message,
	})

	if err != nil {
		app.logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	if app.config.env == "development" {
		trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
		app.logger.Output(2, trace)
	}
	app.logError(r, err)

	message := "the server encountered a problem and could not process your request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	app.errorResponse(w, r, http.StatusNotFound, message)
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.errorResponse(w, r, http.StatusBadRequest, err.Error())
}