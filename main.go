package main

import (
	"net/http"
	"os"
	// "log"
	"fmt"
	log "github.com/sirupsen/logrus"
)

// TODO: add proper error handling for unreachable cat api call
// TODO: write tests for the testing package https://golang.org/pkg/testing/

// set up our logger...maybe better in its own go file.  convention?
// TODO: write a generic logging func for log.WithFields
func init() {
	log.SetFormatter(&log.JSONFormatter{})
	//create log file with read/write permissions
	f, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		// log.Fatal(err)
		log.WithFields(log.Fields{
			"Log Format": "JSON",
			"File":       "log.txt",
		}).Warning("***** Failed To Start Logging *****")
	}
	//defer to close
	//defer f.Close()  // and later reopen it?
	//set output of logs to file log.txt
	log.SetOutput(f)
	//test case
	// log.Info("***** Start Logging *****")
	log.WithFields(log.Fields{
		"Log Format": "JSON",
		"File":       "log.txt", // what's the better way to get this?
	}).Info("***** Start Logging *****")
}

func main() {
	// set up our router, create a new instance of one
	router := NewRouter()

	// let it listen on port :8080 or log the error on failure
	log.Fatal(http.ListenAndServe(":8080", router))
	log.WithFields(log.Fields{
		"Router": router,
		"Port":   ":8080",
	}).Info("Router listening and serving on port 8080")
}

// a general error handling function
func HandleError(err error) {
	if err != nil {
		log.Fatal("Called HandleError", err)
		fmt.Println("fmt Called HandleError", err)
		panic(err)
	}
}
