package main

import (
	"net/http"

	"github.com/ryanpback/raspberry_a_pi/httpd/handlers"
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
	routes := []handlerInfo{
		{
			route:       "/",
			handlerFunc: handlers.Index,
			corsMethods: []string{"GET", "OPTIONS"},
		},
		{
			route:       "/login",
			handlerFunc: handlers.UsersLogin,
			corsMethods: []string{"POST", "OPTIONS"},
		},
		{
			route:       "/users/{id: [0-9]+}",
			handlerFunc: handlers.UsersShow,
			corsMethods: []string{"GET", "OPTIONS"},
		},
		{
			route:       "/users",
			handlerFunc: handlers.UsersCreate,
			corsMethods: []string{"POST", "OPTIONS"},
		},
		{
			route:       "/users/{id:[0-9]+}",
			handlerFunc: handlers.UsersEdit,
			corsMethods: []string{"PATCH", "OPTIONS"},
		},
	}

	return routes
}
