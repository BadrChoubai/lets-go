package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}

// UserModel - type which wraps a database connection pool
type UserModel struct {
	DB *sql.DB
}

func (model *UserModel) Insert(name, email, password string) error {
	return nil
}

func (model *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

func (model *UserModel) Exists(id int) (bool, error) {
	return false, nil
}
