package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/aerytheia/website/internal/api"
	"github.com/aerytheia/website/internal/app"
	"github.com/lmittmann/tint"
)

func main() {
	// setup slog (with colours :o)
	w := os.Stderr

	slog.SetDefault(slog.New(
		tint.NewHandler(w, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.Kitchen,
		}),
	))

	// try and create the app
	app, err := app.NewApp()
	if err != nil {
		slog.Error("failed to create app", "err", err)
		os.Exit(1)
	}

	// create the mux and register the routes
	mux := http.NewServeMux()
	api.RegisterRoutes(mux, app)

	slog.Info("starting server", "port", 8080)
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		slog.Error("failed to start server", "err", err)
	}
}
