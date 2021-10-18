/*
* links.go
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
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/matteoggl/fireside/internal/data"
	"github.com/matteoggl/fireside/internal/validator"
)

func (app *application) homeHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string
		Tags  []string
		data.Filters
	}

	v := validator.New()

	qs := r.URL.Query()

	input.Title = app.readString(qs, "title", "")
	input.Tags = app.readCSV(qs, "tags", []string{})
	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 10, v)
	input.Filters.Sort = app.readString(qs, "sort", "-created_at")
	input.Filters.SortSafelist = []string{"created_at", "-created_at"}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.serverErrorResponse(w, r, errors.New("validation error"))
		return
	}

	links, pagedata, err := app.models.Links.All(input.Title, input.Tags, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

	err = app.inertia.Render(w, r, "Index", map[string]interface{}{
		"links":    links,
		"pagedata": pagedata,
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
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.inertia.Render(w, r, "Show", map[string]interface{}{
		"link": link,
	})

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func (app *application) createHandler(w http.ResponseWriter, r *http.Request) {
	err := app.inertia.Render(w, r, "Create", nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) storeHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string `json:"title"`
		Type    string `json:"type"`
		URL     string `json:"url"`
		Content string `json:"content"`
		Tags    string
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	var tags []string
	if input.Tags != "" {
		toTrim := strings.Split(input.Tags, ",")
		for _, t := range toTrim {
			tags = append(tags, strings.TrimSpace(t))
		}
	}
	tags = append(tags, input.Type)

	if input.Type == "text" {
		input.URL = ""
	} else {
		input.Content = ""
	}

	link := &data.Link{
		Title:   input.Title,
		Type:    input.Type,
		URL:     sql.NullString{String: input.URL, Valid: input.URL != ""},
		Content: sql.NullString{String: input.Content, Valid: input.Content != ""},
		Tags:    tags,
	}

	v := validator.New()
	if data.ValidateLink(v, link); !v.Valid() {
		app.inertia.Render(w, r, "Create", map[string]interface{}{
			"errors": v.Errors,
		})
		return
	}

	err = app.models.Links.Insert(link)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	link.FormatCreatedAt()

	w.Header().Set("Location", fmt.Sprintf("/links/%d", link.ID))
	w.WriteHeader(http.StatusSeeOther)
}

func (app *application) destroyHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.redIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.Links.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusSeeOther)
}
