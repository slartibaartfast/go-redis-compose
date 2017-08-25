package main

import (
	  "encoding/json"
    // s"fmt"
    "net/http"
	  "io"
	  "io/ioutil"
	  // s"strconv"
	  // s"time"

    mux "github.com/julienschmidt/httprouter"
)


func CatShow(w http.ResponseWriter, r *http.Request, p mux.Params) {

        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(http.StatusOK)

        // var catshow Cat
        catshow := FindAll()

        if err := json.NewEncoder(w).Encode(catshow); err != nil {
                panic(err)
        }
}


func CatCreate(w http.ResponseWriter, r *http.Request, _ mux.Params) {
        //TODO: convert the incoming xml to json
      	var cat Cat

      	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
      	HandleError(err)

      	if err := r.Body.Close(); err != nil {
      	        panic(err)
      	}

      	// Save JSON to Cat struct
      	if err := json.Unmarshal(body, &cat); err != nil {

      	        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
      	        w.WriteHeader(422)

      	        if err := json.NewEncoder(w).Encode(err); err != nil {
      	                panic(err)
      	        }
      	}

      	CreateCat(cat)
              w.Header().Set("Content-Type", "application/json; charset=UTF-8")
              w.WriteHeader(http.StatusCreated)
}


func CatHistory(w http.ResponseWriter, r *http.Request, _ mux.Params) {

        var cats Cats

        cats = FindAll()

        if err := json.NewEncoder(w).Encode(cats); err != nil {
                panic(err)
        }
}
