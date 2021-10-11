/*
* users.go
* Copyright (C) <2021>  <Matteo Guglielmetti>
*
* This program is free software: you can redistribute it and/or modify
* it under the terms of the GNU Affero General Public License as published
* by the Free Software Foundation, either version 3 of the License, or
* (at your option) any later version.
*
* This program is distributed in the hope that it will be useful,
* but WITHOUT ANY WARRANTY; without even the implied warranty of
* MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
* GNU Affero General Public License for more details.
*
* You should have received a copy of the GNU Affero General Public License
* along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package main

import (
	"errors"
	"net/http"

	"github.com/matteoggl/fireside/internal/data"
	"github.com/matteoggl/fireside/internal/validator"
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
