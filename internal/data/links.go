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

package data

import (
	"database/sql"
	"errors"
	"time"

	"github.com/lib/pq"
	"github.com/matteoggl/fireside/internal/validator"
)

type Link struct {
	ID        int64
	Title     string
	Type      string
	URL       sql.NullString
	Content   sql.NullString
	Likes     int64
	Tags      []string
	CreatedAt time.Time
}

func ValidateLink(v *validator.Validator, link *Link) {
	v.Check(link.Title != "", "title", "Title must be provided")
	v.Check(link.Type == "link" || link.Type == "text", "type", "Unspecified type")
	if link.Type == "link" {
		v.Check(link.URL.Valid, "url", "URL must be specified")
	}
	if link.Type == "text" {
		v.Check(link.Content.Valid, "content", "Content must be specified")
	}
}

type LinkModel struct {
	DB *sql.DB
}

func (l LinkModel) All() ([]*Link, error) {
	stmt := `SELECT id, title, type, url, content, likes, tags, created_at
	FROM links
	ORDER BY created_at DESC`

	rows, err := l.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	links := []*Link{}
	for rows.Next() {
		l := &Link{}
		err = rows.Scan(
			&l.ID,
			&l.Title,
			&l.Type,
			&l.URL,
			&l.Content,
			&l.Likes,
			pq.Array(&l.Tags),
			&l.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		links = append(links, l)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return links, nil
}

func (l LinkModel) Insert(link *Link) error {
	stmt := `INSERT INTO links (title, type, url, content, tags)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, created_at`

	return l.DB.QueryRow(stmt, &link.Title, &link.Type, &link.URL, &link.Content, pq.Array(&link.Tags)).Scan(&link.ID, &link.CreatedAt)
}

func (l LinkModel) Get(id int64) (*Link, error) {
	stmt := `SELECT id, title, type, url, content, likes, tags, created_at
	FROM links
	WHERE id = $1`

	var link Link
	err := l.DB.QueryRow(stmt, id).Scan(
		&link.ID,
		&link.Title,
		&link.Type,
		&link.URL,
		&link.Content,
		&link.Likes,
		pq.Array(&link.Tags),
		&link.CreatedAt,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &link, nil
}

func (l LinkModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	stmt := `
	DELETE FROM links
	WHERE id = $1`

	res, err := l.DB.Exec(stmt, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}

func (l LinkModel) ByTag(tag string) ([]*Link, error) {
	stmt := `SELECT id, title, type, url, content, likes, tags, created_at
	FROM links
	WHERE $1 = ANY (tags)
	ORDER BY created_at DESC`

	rows, err := l.DB.Query(stmt, tag)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	links := []*Link{}
	for rows.Next() {
		l := &Link{}
		err = rows.Scan(
			&l.ID,
			&l.Title,
			&l.Type,
			&l.URL,
			&l.Content,
			&l.Likes,
			pq.Array(&l.Tags),
			&l.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		links = append(links, l)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return links, nil
}
