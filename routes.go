package main

import (
	"github.com/harrydrippin/go-quick-link/server"
)

var routes = server.Routes{
	server.Route{
		Name:        "Health",
		Method:      "GET",
		Path:        "/health",
		HandlerFunc: HandlerHealthCheck,
	},
	server.Route{
		Name:        "Go",
		Method:      "GET",
		PathPrefix:  "/go",
		HandlerFunc: HandlerGo,
	},
	server.Route{
		Name:        "Get Quick Links",
		Method:      "GET",
		Path:        "/manage",
		HandlerFunc: HandlerGetQuickLinks,
	},
	server.Route{
		Name:        "Set Quick Link",
		Method:      "POST",
		Path:        "/manage",
		HandlerFunc: HandlerSetQuickLinks,
	},
	server.Route{
		Name:        "Delete Quick Link",
		Method:      "DELETE",
		Path:        "/manage",
		HandlerFunc: HandlerDeleteQuickLinks,
	},
	// TODO(@harry): Support Batch Import / Export by JSON
}
