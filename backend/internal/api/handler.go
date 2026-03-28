package api

import (
	"net/http"

	"github.com/aerytheia/website/internal/app"
)

func RegisterRoutes(mux *http.ServeMux, a *app.App) {
	mux.HandleFunc("POST /api/login", func(w http.ResponseWriter, r *http.Request) {
		PostLogin(w, r, a)
	})
}
