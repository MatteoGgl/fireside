package data

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
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

type LinkModel struct {
	DB *sql.DB
}

func (l *LinkModel) All() ([]*Link, error) {
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

func (l *LinkModel) Get(id int64) (*Link, error) {
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
		return nil, err
	}

	return &link, nil
}

func (l *LinkModel) ByTag(tag string) ([]*Link, error) {
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
