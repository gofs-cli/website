package frontend

import (
	"net/http"

	"github.com/a-h/templ"

	"github.com/gofs-cli/website/internal/ui"
	"github.com/gofs-cli/website/internal/ui/pages/docs"
)

func NoapiPage() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(ui.Index(docs.Docs(Noapi(), "/docs/frontend/noapi"))).ServeHTTP(w, r)
	})
}

func AddPage() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(ui.Index(docs.Docs(Page(), "/docs/frontend/page"))).ServeHTTP(w, r)
	})
}

func HandlerPage() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(ui.Index(docs.Docs(Handler(), "/docs/frontend/handler"))).ServeHTTP(w, r)
	})
}

func RoutingPage() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(ui.Index(docs.Docs(Routing(), "/docs/frontend/routing"))).ServeHTTP(w, r)
	})
}
