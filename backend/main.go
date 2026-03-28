package main

import (
	"net/http"

	"github.com/aerytheia/website/internal/app"
)

func main() {
	// setup slog

	// make mux
	// serve routes

	app := app.NewApp()
	mux := http.NewServeMux()

}
