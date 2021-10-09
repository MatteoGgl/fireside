package main

import (
	"errors"
	"net/http"

	"github.com/matteoggl/linki/internal/data"
	"github.com/matteoggl/linki/internal/validator"
)

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
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user := &data.User{
		Email:          input.Email,
		HashedPassword: []byte(input.Password),
	}

	v := validator.New()
	if data.ValidateUser(v, user); !v.Valid() {
		app.inertia.Render(w, r, "Login", map[string]interface{}{
			"errors": v.Errors,
		})
		return
	}

	id, err := app.models.Users.Authenticate(input.Email, input.Password)
	if err != nil {
		if errors.Is(err, data.ErrInvalidCredentials) {
			v.AddError("generic", "Email or Password is incorrect")
			app.inertia.Render(w, r, "Login", map[string]interface{}{
				"errors": v.Errors,
			})
		} else {
			app.badRequestResponse(w, r, err)
		}
		return
	}

	app.session.Put(r, "authenticatedUserID", id)

	path := app.session.PopString(r, "redirectPathAfterLogin")
	if path != "" {
		http.Redirect(w, r, path, http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) logoutHandler(w http.ResponseWriter, r *http.Request) {
	app.session.Remove(r, "authenticatedUserID")

	http.Redirect(w, r, "/", http.StatusSeeOther)
}