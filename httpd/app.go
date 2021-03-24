package main

import (
	"net/http"

	h "github.com/ryanpback/raspberry_a_pi/helpers"
	"github.com/ryanpback/raspberry_a_pi/httpd/middleware"

	"github.com/gorilla/mux"
)

func runApp() {
	r := mux.NewRouter()

	for _, handlerOpts := range openRoutes() {
		r.HandleFunc(
			handlerOpts.route,
			middleware.HandleCors(handlerOpts.handlerFunc),
		).
			Methods(handlerOpts.corsMethods...)
	}

	for _, handlerOpts := range paramRoutes() {
		r.HandleFunc(
			handlerOpts.route,
			middleware.HandleCors(handlerOpts.handlerFunc),
		).
			Queries(convertQueries(handlerOpts.queries)...).
			Methods(handlerOpts.corsMethods...)
	}

	h.LogInfo("Starting web api at port " + appConfig.AppPort)

	h.LogFatal(http.ListenAndServe(":"+appConfig.AppPort, r))
}

func convertQueries(q query) []string {
	pairs := make([]string, 0, len(q)*2)
	for key, value := range q {
		pairs = append(pairs, key, value)
	}

	return pairs
}
