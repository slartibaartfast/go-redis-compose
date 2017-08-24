package main

import (
        mux "github.com/julienschmidt/httprouter"
)

// Create a router instance
//func NewRouter *mux.Router {
func NewRouter() *mux.Router {

        router := mux.New()

        for _, route := range routes {

                router.Handle(route.Method, route.Pattern, route.Handle)

        }

        return router
}
