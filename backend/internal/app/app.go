package app

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type App struct {
	db *sql.DB
}

func NewApp() (*App, error) {
	db, err := sql.Open("sqlite3", "./aerytheia.db")
	if err != nil {
		return nil, fmt.Errorf("unable to open database: %w", err)
	}

	return &App{
		db: db,
	}, nil
}
