package server

import (
	"net/http"

	"github.com/gofs-cli/website/internal/server/assets"
	"github.com/gofs-cli/website/internal/server/handlers"
	"github.com/gofs-cli/website/internal/server/telemetry/metrics"
	"github.com/gofs-cli/website/internal/ui/pages"
	"github.com/gofs-cli/website/internal/ui/pages/blogs"
	"github.com/gofs-cli/website/internal/ui/pages/docs"
	backend "github.com/gofs-cli/website/internal/ui/pages/docs/Backend"
)

func (s *Server) Routes() {
	// filserver route for assets
	assetMux := http.NewServeMux()
	assetMux.Handle("GET /{path...}", http.StripPrefix("/assets/", handlers.NewHashedAssets(assets.FS)))
	s.r.Handle("GET /assets/{path...}", s.assetsMiddlewares(assetMux))

	// handlers for normal routes with all general middleware
	routesMux := http.NewServeMux()
	routesMux.Handle("GET /{$}", pages.HomePage())
	// docs
	routesMux.Handle("GET /docs", docs.PrereqsPage())
	routesMux.Handle("GET /docs/prereqs", docs.PrereqsPage())
	routesMux.Handle("GET /docs/start", docs.StartPage())
	routesMux.Handle("GET /docs/backend/module", backend.ModulePage())
	routesMux.Handle("GET /docs/backend/persistence", backend.PersistencePage())
	routesMux.Handle("GET /docs/backend/codegen", backend.CodegenPage())
	// blogs
	routesMux.Handle("GET /blogs", blogs.IndexPage())
	routesMux.Handle("GET /blogs/simplicity", blogs.SimplicyBlog())
	routesMux.Handle("GET /blogs/libraries", blogs.LibrariesBlog())
	routesMux.Handle("GET /blogs/nosql", blogs.NosqlBlog())
	routesMux.Handle("GET /blogs/evergreen", blogs.EvergreenBlog())

	s.r.Handle("/", s.routeMiddlewares(routesMux))

	s.srv.Handler = metrics.Expose(s.r)
}
