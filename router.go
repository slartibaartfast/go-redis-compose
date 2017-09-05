package main

import (
	"fmt"
	mux "github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

// create a router instance
func NewRouter() *mux.Router {
	router := mux.New()
	fmt.Println("fmt new router ", router)
	log.Println("log new router")

	// loop through our routes and save them
	for _, route := range routes {
		router.Handle(route.Method, route.Pattern, route.Handle)
	}
	fmt.Println("fmt return router ", router)
	log.Println("log return router", router)
	return router
}
