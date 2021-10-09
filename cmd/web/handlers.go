package main

import "net/http"

func (app *application) homeHandler(w http.ResponseWriter, r *http.Request) {
	links, err := app.models.Links.All()
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

	err = app.inertia.Render(w, r, "Links", map[string]interface{}{
		"links": links,
	})

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func (app *application) showHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.redIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	link, err := app.models.Links.Get(id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.inertia.Render(w, r, "Link", map[string]interface{}{
		"link": link,
	})

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func (app *application) byTagHandler(w http.ResponseWriter, r *http.Request) {
	tag := app.readString(r, "tag", "all")

	links, err := app.models.Links.ByTag(tag)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

	err = app.inertia.Render(w, r, "Links", map[string]interface{}{
		"links": links,
	})

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func (app *application) loginFormHandler(w http.ResponseWriter, r *http.Request) {
	err := app.inertia.Render(w, r, "Login", nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func (app *application) loginHandler(w http.ResponseWriter, r *http.Request) {

}
