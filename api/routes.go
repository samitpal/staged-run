package api

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"RunIndex",
		"GET",
		"/run",
		RunIndex,
	},
	Route{
		"RunCreate",
		"POST",
		"/run",
		RunCreate,
	},
	Route{
		"RunStatus",
		"GET",
		"/run/{runId}",
		RunStatus,
	},
}
