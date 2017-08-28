package main

import (
	mux "github.com/julienschmidt/httprouter"
)

// create our route structure
type Route struct {
	Name    string
	Method  string
	Pattern string
	Handle  mux.Handle // handler function, the httprouter package as mux
}

// slice Route
type Routes []Route

// define the url routes to our endpoints
var routes = Routes{
	// GET a picture of a cat from http://thecatapi.com/
	Route{
		"CatShow",
		"GET",
		"/cat",
		CatShow,
	},
	// GET all queries to /cat endpoint
	Route{
		"CatHistory",
		"GET",
		"/history",
		CatHistory,
	},
	// POST to create a cat record
	// not using this atm
	Route{
		"CatCreate",
		"POST",
		"/cat",
		CatCreate,
	},
}
