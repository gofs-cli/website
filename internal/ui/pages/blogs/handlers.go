package blogs

import (
	"net/http"

	"github.com/a-h/templ"

	"github.com/gofs-cli/website/internal/ui"
)

func IndexPage() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(ui.Index(Blogs(Index()))).ServeHTTP(w, r)
	})
}

func SimplicyBlog() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(ui.Index(Blogs(Simplcity()))).ServeHTTP(w, r)
	})
}

func LibrariesBlog() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(ui.Index(Blogs(Libraries()))).ServeHTTP(w, r)
	})
}

func NosqlBlog() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(ui.Index(Blogs(Nosql()))).ServeHTTP(w, r)
	})
}

func EvergreenBlog() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(ui.Index(Blogs(Evergreen()))).ServeHTTP(w, r)
	})
}
