package server

import (
	"net/http"

	"github.com/gofs-cli/website/internal/server/assets"
	"github.com/gofs-cli/website/internal/server/handlers"
	"github.com/gofs-cli/website/internal/ui/pages"
	"github.com/gofs-cli/website/internal/ui/pages/blogs"
	backend "github.com/gofs-cli/website/internal/ui/pages/docs/Backend"
	frontend "github.com/gofs-cli/website/internal/ui/pages/docs/Frontend"
	prereqs "github.com/gofs-cli/website/internal/ui/pages/docs/Prereqs"
	start "github.com/gofs-cli/website/internal/ui/pages/docs/Start"
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
	routesMux.Handle("GET /docs", prereqs.PrereqsPage())
	routesMux.Handle("GET /docs/prereqs", prereqs.PrereqsPage())
	routesMux.Handle("GET /docs/prereqs/database", prereqs.DatabasePage())
	routesMux.Handle("GET /docs/prereqs/npm", prereqs.NpmPage())
	routesMux.Handle("GET /docs/start", start.StartPage())
	routesMux.Handle("GET /docs/start/tailwind", start.TailwindPage())
	routesMux.Handle("GET /docs/backend/module", backend.ModulePage())
	routesMux.Handle("GET /docs/backend/persistence", backend.PersistencePage())
	routesMux.Handle("GET /docs/backend/codegen", backend.CodegenPage())
	routesMux.Handle("GET /docs/frontend/noapi", frontend.NoapiPage())
	routesMux.Handle("GET /docs/frontend/page", frontend.AddPage())
	routesMux.Handle("GET /docs/frontend/handler", frontend.HandlerPage())
	routesMux.Handle("GET /docs/frontend/routing", frontend.RoutingPage())
	// blogs
	routesMux.Handle("GET /blogs", blogs.IndexPage())
	routesMux.Handle("GET /blogs/simplicity", blogs.SimplicyBlog())
	routesMux.Handle("GET /blogs/libraries", blogs.LibrariesBlog())
	routesMux.Handle("GET /blogs/nosql", blogs.NosqlBlog())
	routesMux.Handle("GET /blogs/evergreen", blogs.EvergreenBlog())

	s.r.Handle("/", s.routeMiddlewares(routesMux))

	s.srv.Handler = s.r
}
