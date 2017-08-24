package main

import (
        mux "github.com/julienschmidt/httprouter"
)


// Create our route
type Route struct {
        Name       string
        Method     string
        Pattern    string
        Handle     mux.Handle   // handler function, the httprouter package as mux
}


// Slice Route
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
        // POST to create a cat record
        Route{
                "CatCreate",
                "POST",
                "/cat",
                CatCreate,
        },
        // GET all queries
        Route{
                "CatHistory",
                "GET",
                "/history",
                CatHistory,
        },
}
