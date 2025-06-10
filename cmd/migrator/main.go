package main

import (
	"database/sql"
	"log"

	"github.com/7ngg/tg-water/internal/config"
	"github.com/pressly/goose/v3"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	cfg := config.MustLoad()
	const op = "cmd.migrator.main"

	db, err := sql.Open("sqlite3", cfg.Storage.Path)
	if err != nil {
		log.Fatalf("%s: %v", op, err)
	}

	if err := goose.SetDialect("sqlite3"); err != nil {
		log.Fatalf("%s: %v", op, err)
	}

	if err := goose.Up(db, cfg.Storage.Migrations); err != nil {
		log.Fatalf("%s: %v", op, err)
	}

}
