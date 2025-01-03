package docs

import (
	"net/http"

	"github.com/a-h/templ"

	"github.com/gofs-cli/website/internal/ui"
)

func StartPage() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(ui.Index(Docs(Start(), "/docs/start"))).ServeHTTP(w, r)
	})
}

func PrereqsPage() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(ui.Index(Docs(Prereqs(), "/docs/prereqs"))).ServeHTTP(w, r)
	})
}
