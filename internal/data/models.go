/*
* models.go
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

package data

import (
	"database/sql"
	"errors"
)

var (
	ErrNoRecord           = errors.New("no matching record found")
	ErrRecordNotFound     = errors.New("record not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type Models struct {
	Links LinkModel
	Users UserModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Links: LinkModel{DB: db},
		Users: UserModel{DB: db},
	}
}
