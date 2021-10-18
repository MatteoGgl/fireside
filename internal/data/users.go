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

package data

import (
	"database/sql"
	"errors"
	"time"

	"github.com/matteoggl/fireside/internal/validator"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrDuplicateEmail = errors.New("duplicate email")
)

type User struct {
	ID             int64
	Email          string
	HashedPassword []byte
	CreatedAt      time.Time
}

func ValidateUser(v *validator.Validator, user *User) {
	v.Check(user.Email != "", "email", "Email must be provided")
	v.Check(validator.Matches(user.Email, validator.EmailRX), "email", "Email must be a valid email address")
	v.Check(len(user.HashedPassword) > 0, "password", "A password must be provided")
}

type UserModel struct {
	DB *sql.DB
}

func (u *UserModel) Insert(email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := `INSERT INTO users (email, hashed_password)
	VALUES ($1, $2)`

	_, err = u.DB.Exec(stmt, email, string(hashedPassword))
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return ErrDuplicateEmail
		default:
			return err
		}
	}

	return nil
}

func (u *UserModel) Authenticate(email, password string) (int, error) {
	var id int
	var hashedPassword []byte

	stmt := `SELECT id, hashed_password
	FROM users
	WHERE email = $1`
	row := u.DB.QueryRow(stmt, email)
	err := row.Scan(&id, &hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, ErrInvalidCredentials
		} else {
			return 0, err
		}
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, ErrInvalidCredentials
		} else {
			return 0, err
		}
	}

	return id, nil
}

func (um *UserModel) Get(id int) (*User, error) {
	u := &User{}

	stmt := `SELECT id, email FROM users WHERE id = $1`
	err := um.DB.QueryRow(stmt, id).Scan(
		&u.ID,
		&u.Email,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrRecordNotFound
		} else {
			return nil, err
		}
	}

	return u, nil
}
