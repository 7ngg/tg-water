package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	DB *Queries
}

func NewConnection(storagePath string) (*Storage, error) {
	const op = "storage.sqlite.New"

	connection, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	db := New(connection)

	return &Storage{DB: db}, nil
}
