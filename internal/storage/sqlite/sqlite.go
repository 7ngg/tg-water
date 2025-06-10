package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *Queries
}

func NewConnection(storagePath string) (*Storage, error) {
	const op = "storage.sqlite.New"

	connection, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	db := New(connection)

	return &Storage{db: db}, nil
}
