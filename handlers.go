package main

import (
	  "encoding/json"
		"encoding/xml"
    "fmt"
    "net/http"
	  "io"
	  "io/ioutil"
		// s"strings"
	  // s"strconv"
	  "time"
    // xj "github.com/basgys/goxml2json"
    mux "github.com/julienschmidt/httprouter"
)


func CatShow(w http.ResponseWriter, r *http.Request, p mux.Params) {

        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(http.StatusOK)

		    //var catshow Cat
				var image Image

				//catshow = FetchCat("http://thecatapi.com/api/images/get?format=xml&results_per_page=1")
				image = FetchCat("http://thecatapi.com/api/images/get?format=xml&results_per_page=1")

				//if err := json.NewEncoder(w).Encode(catshow); err != nil {
        if err := json.NewEncoder(w).Encode(image); err != nil {
                panic(err)
        }

				//CreateCat(catshow)
				CreateCat(image.Image)
              w.Header().Set("Content-Type", "application/json; charset=UTF-8")
              w.WriteHeader(http.StatusCreated)

}


//func FetchCat(url string) Cat {
func FetchCat(url string) Image {
	var myClient = &http.Client{Timeout: 10 * time.Second}

  //apiKey := "YouDontNeedToKnow"
	//var query string
	//query := url + apiKey
  //r, err := myClient.Get()
	r, err := myClient.Get(url)
	if err != nil {
	//		return err
	      fmt.Println(err)
	}
	defer r.Body.Close()

  var body []byte
	body, err = ioutil.ReadAll(r.Body)
	//xml := strings.NewReader(string(body))
	fmt.Printf("the body: %s\n", body)

  // unpack the response into our xml structs formats
	var dict xmlResponse
	xml.Unmarshal(body, &dict)
	//fmt.Printf("unmarshalled xml: %s\n", Images)
	fmt.Println("dict.Images: [%s]\n", dict.Images)

	// convert to json
	//var oneCat Cat
  //var allCats []Cat  // struct Cats
	var oneCat Image
	//var allCats Images


	for _, value := range dict.Images {
					//oneCat.Id = value.Id
					oneCat.Image.Id = value.Id
					oneCat.Image.Url = value.Url
					oneCat.Image.Source_Url = value.Source_Url

					//allCats = append(allCats, oneCat)
					//allCats = allCats.oneCat
	}

	//jsonData, err := json.Marshal(allCats)

	//if err != nil {
	//				fmt.Println(err)
	//				//os.Exit(1)
	//}

  //fmt.Println(string(jsonData))

	//fmt.Printf("JSON for output: \n%s\n",allCats)

	return oneCat
}

// called from history endpoint
func CatHistory(w http.ResponseWriter, r *http.Request, _ mux.Params) {

        //var cats Cats
				var cats Images
				//var images Images

        cats = FindAll()
				//images.Images = cats

        if err := json.NewEncoder(w).Encode(cats); err != nil {
				//if err := json.NewEncoder(w).Encode(images.Images); err != nil {
                panic(err)
        }
}


// this gets called on post, which we are not worried about atm
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
