package app

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

type App struct {
	db *sql.DB
}

func NewApp() (*App, error) {
	db, err := sql.Open("sqlite", "./aerytheia.db")
	if err != nil {
		return nil, fmt.Errorf("unable to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("unable to ping database: %w", err)
	}

	return &App{
		db: db,
	}, nil
}
