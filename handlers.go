package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	mux "github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	// "log"
	log "github.com/sirupsen/logrus"
)

// show json from thecatapi.com representing an individual cat
func CatShow(w http.ResponseWriter, r *http.Request, p mux.Params) {
	// set our header to json
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// we will store data in an Image structure
	var image Image

	// load our struct with data retrieved by the FetchCat function
	// image = FetchCat(os.Getenv("CAT_URL"))
	image = FetchCat()
	if err := json.NewEncoder(w).Encode(image); err != nil {
		HandleError(err)
	}

	// call a db function to load our cat image data into redis and print result
	CreateCat(image.Image)

	// write a response header with a created status value
	w.WriteHeader(http.StatusCreated)
}

// pass in a url and recieve data back per the Image struct
// TODO: separate some of this into other functions
func FetchCat() Image {
	// make an http client and set our connection timeout
	client := &http.Client{Timeout: 10 * time.Second}
	//client, err := &http.Client{Timeout: 10 * time.Second}
	//if err != nil {
	//		HandleError(err)
	//}
	fmt.Println("fmt fetch ")
	log.Println("log fetch ")

	// set up our url string
	baseUrl := os.Getenv("CAT_URL")
	action := "images/get?"

	// read the Docker Secrets file "api_key", and check for errors
	b, err := ioutil.ReadFile(os.Getenv("CAT_API_KEY_PASSWORD_FILE"))
  if err != nil {
      HandleError(err)
  }
	log.Println("secrets body ", b)
	log.Println("err ", err)

	// convert content of secrets file to a string
  apiKey := "api_key=" + string(b)
	log.Println("api_key ", apiKey)

  // set up our new request with the beggining of the url string
	req, err := http.NewRequest("GET", baseUrl + action, nil)
	if err != nil {
		HandleError(err)
	}

	// add parameters to our url and encode them
	query := req.URL.Query()
	query.Add("api_key", apiKey)
	query.Add("format", "xml")
	query.Add("results_per_page", "1")
	req.URL.RawQuery = query.Encode()

	// add a header to be polite and possibley negate cross site request forgery
	req.Header.Set("X-Requested-With", "XMLHttpRequest")

	log.WithFields(log.Fields{
	  "api key": apiKey,
	  "query": query,
	  "request": req,
	}).Info("Sending this")

	// open a connection, make our request and defer closure of the connection
	resp, err := client.Do(req)
	if err != nil {
		HandleError(err)
	}
	defer resp.Body.Close()

	// read the response into a variable
	var body []byte
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		HandleError(err)
	}

	// unpack the response into our xmlResponse structs format
	var dict xmlResponse
	xml.Unmarshal(body, &dict)

	// loop through the xml, convert to json by copying values to Image struct
	// note: it might be better to do this with a mapping struct
	var oneCat Image
	for _, value := range dict.Images {
		oneCat.Image.Id = value.Id
		fmt.Printf("the catid: %s\n", oneCat.Image.Id)
		oneCat.Image.Url = value.Url
		oneCat.Image.Source_Url = value.Source_Url
	}

	// return the json as an Image
	return oneCat
}

// called from /history endpoint, returns all json strings from requests to /cat
func CatHistory(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	// set Images to store our data
	var cats Images

	// run the db fucntion to return the saved json
	cats = FindAll()

	// encode the data in our Images struct as json string and write it out
	if err := json.NewEncoder(w).Encode(cats); err != nil {
		HandleError(err)
	}
}

// this gets called on post, which we are not worried about atm
// TODO: finsih and test this
func CatCreate(w http.ResponseWriter, r *http.Request, _ mux.Params) {
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
