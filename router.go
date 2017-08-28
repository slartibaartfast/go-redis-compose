package main

import (
	mux "github.com/julienschmidt/httprouter"
)

// create a router instance
func NewRouter() *mux.Router {
	router := mux.New()

	// loop through our routes and save them
	for _, route := range routes {
		router.Handle(route.Method, route.Pattern, route.Handle)
	}
	return router
}
