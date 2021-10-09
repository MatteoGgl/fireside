package data

import (
	"database/sql"
	"errors"
	"time"

	"github.com/matteoggl/linki/internal/validator"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             int64
	Email          string
	HashedPassword []byte
	CreatedAt      time.Time
}

func ValidateUser(v *validator.Validator, user *User) {
	v.Check(user.Email != "", "email", "Email must be provided")
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
		return err
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
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return u, nil
}
