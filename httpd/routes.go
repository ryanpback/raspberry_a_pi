package main

import (
	"net/http"
)

type query map[string]string

type handlerInfo struct {
	route       string
	queries     query
	handlerFunc func(http.ResponseWriter, *http.Request)
	corsMethods []string
}

func paramRoutes() []handlerInfo {
	var routes = []handlerInfo{
		{
			// route: "/some_route",
			// queries: query{
			// 	"key": "{value}",
			// },
			// handlerFunc: handlers.SomeMethod,
			// corsMethods: []string{http.HttpVerb},
		},
	}

	return routes
}

func openRoutes() []handlerInfo {
	var routes []handlerInfo

	routes = []handlerInfo{
		{
			// route:       "/some_route",
			// handlerFunc: handlers.SomeMethod,
			// corsMethods: []string{http.HttpVerb},
		},
	}

	return routes
}
