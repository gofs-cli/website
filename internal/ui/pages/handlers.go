package pages

import (
	"net/http"

	"github.com/a-h/templ"

	"github.com/gofs-cli/website/internal/ui"
)

func HomePage() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(ui.Index(Home())).ServeHTTP(w, r)
	})
}
