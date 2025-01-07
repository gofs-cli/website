package pages

import (
	"slices"

	"github.com/gofs-cli/website/internal/ui/components"
)

var LeftMenu = []components.NavItem{
	{
		Text: "Prerequisites",
		Path: "/docs/prereqs",
		Children: []components.NavItem{
			{
				Text: "Database",
				Path: "/docs/prereqs/database",
			},
			{
				Text: "Tailwind",
				Path: "/docs/prereqs/tailwind",
			},
		},
	},
	{
		Text: "Start a project",
		Path: "/docs/start",
	},
	{
		Text: "Backend",
		Children: []components.NavItem{
			{
				Text: "Adding a module",
				Path: "/docs/backend/module",
			},
			{
				Text: "Persisting data",
				Path: "/docs/backend/persistence",
			},
			{
				Text: "Codegen",
				Path: "/docs/backend/codegen",
			},
		},
	},
	{
		Text: "Frontend",
		Children: []components.NavItem{
			{
				Text: "No API",
				Path: "/docs/frontend/noapi",
			},
			{
				Text: "Adding a page",
				Path: "/docs/frontend/page",
			},
			{
				Text: "Routing",
				Path: "/docs/frontend/routing",
			},
			{
				Text: "Adding a handler",
				Path: "/docs/frontend/handler",
			},
		},
	},
}

var HamburgerMenu = slices.Concat(LeftMenu, []components.NavItem{
	{
		Text: "Components",
		Path: "/components",
	},
	{
		Text: "Blogs",
		Path: "/blogs",
	},
})
