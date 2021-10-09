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
