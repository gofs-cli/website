package start

import (
	"net/http"

	"github.com/a-h/templ"

	"github.com/gofs-cli/website/internal/ui"
	"github.com/gofs-cli/website/internal/ui/pages/docs"
)

func StartPage() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(ui.Index(docs.Docs(Start(), "/docs/start"))).ServeHTTP(w, r)
	})
}

func TailwindPage() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(ui.Index(docs.Docs(Tailwind(), "/docs/start/tailwind"))).ServeHTTP(w, r)
	})
}
