package main

import (
	"log"
	"net/http"
)

func main() {
	// set up our router, create a new instance of one
	router := NewRouter()

	// let it listen on port :8080 or log the error on failure
	log.Fatal(http.ListenAndServe(":8080", router))
}

// a general error handling function
func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
