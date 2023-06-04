package models

import (
	"database/sql"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// SnippetModel Think of this as a service layer or data access layer
type SnippetModel struct {
	DB *sql.DB
}

func (model *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	return 0, nil
}

func (model *SnippetModel) Get(id int) (*Snippet, error) {
	return nil, nil
}
func (model *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
