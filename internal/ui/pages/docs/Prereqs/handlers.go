package prereqs

import (
	"net/http"

	"github.com/a-h/templ"

	"github.com/gofs-cli/website/internal/ui"
	"github.com/gofs-cli/website/internal/ui/pages/docs"
)

func PrereqsPage() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(ui.Index(docs.Docs(Prereqs(), "/docs/prereqs"))).ServeHTTP(w, r)
	})
}

func DatabasePage() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(ui.Index(docs.Docs(Database(), "/docs/prereqs/database"))).ServeHTTP(w, r)
	})
}

func TailwindPage() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(ui.Index(docs.Docs(Tailwind(), "/docs/prereqs/tailwind"))).ServeHTTP(w, r)
	})
}
