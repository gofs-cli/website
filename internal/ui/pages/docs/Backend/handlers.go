package backend

import (
	"net/http"

	"github.com/a-h/templ"

	"github.com/gofs-cli/website/internal/ui"
	"github.com/gofs-cli/website/internal/ui/pages/docs"
)

func ModulePage() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(ui.Index(docs.Docs(Module(), "/docs/backend/module"))).ServeHTTP(w, r)
	})
}

func PersistencePage() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(ui.Index(docs.Docs(Persistence(), "/docs/backend/persistence"))).ServeHTTP(w, r)
	})
}

func CodegenPage() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(ui.Index(docs.Docs(Codegen(), "/docs/backend/codegen"))).ServeHTTP(w, r)
	})
}
