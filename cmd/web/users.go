package main

import (
	"errors"
	"net/http"

	"github.com/matteoggl/linki/internal/data"
	"github.com/matteoggl/linki/internal/validator"
)


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