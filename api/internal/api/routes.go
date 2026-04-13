package api

import (
	"io"
	"net/http"
	"os/exec"

	"github.com/aerytheia/website/internal/app"
	"github.com/creack/pty"
)

func GetFastfetch(w http.ResponseWriter, r *http.Request, a *app.App) {
	cmd := exec.Command("fastfetch")

	ptmx, err := pty.Start(cmd)
	if err != nil {
		http.Error(w, "failed to start pty", 500)
		return
	}
	defer ptmx.Close()

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	// stream output directly
	io.Copy(w, ptmx)
}

func PostGuestbook(w http.ResponseWriter, r *http.Request, a *app.App) {
	
}

func PostLogin(w http.ResponseWriter, r *http.Request, a *app.App) {
	
}
