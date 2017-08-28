package main

import (
  "encoding/xml"
  // s"encoding/json"
)

//TODO: you might only need the cat struct.  Just add all queries to the external
// cat api to it, use it to return them one at the /cat endpoints
// FindAll to give them the history endpoint

// cats structure
type Cat struct {
        Id           string  `json:"id"`          // the id from the cat api (used to be our internal int id)
        Url          string  `json:"url"`         // the original url (tumblr, imgur, redit, ...)
        Source_Url   string  `json:"source_url"`  // url of the image on the cat api
}

type Image struct {
    Image  Cat  `json:"image"`
}

type Images struct {
    Images []Cat `json:"images"`
    //Images Cats `json:"images"`
}

//type jsonImages struct {
//
//}

// to match the cat api xml output
// not calling this anymore ?
type xmlImages struct {
  Response   xml.CharData `xml:"response"`
      Data   xml.CharData    `xml:"data"`
          Images  xml.CharData  `xml:"images>image"`
              //Image []struct {
              Image xml.CharData `xml:image`
                  Url    xml.CharData `xml:"url"`
                  Id     xml.CharData    `xml:"id"`
                  Source_Url  xml.CharData  `xml:"Source_Url"`
              //} `xml:"image"`

}


type xmlResponse struct {
	Response   xml.Name  `xml:"response"`
  //FirstImage string    `xml:"data>images>image"`
	Images []xmlImage `xml:"data>images>image"`
}


type xmlImage struct {
      Image xml.Name  `xml:"image"`
      Url    string `xml:"url"`
      Id     string    `xml:"id"`
      Source_Url  string  `xml:"source_url"`
}


type Cats []Cat


// history structure, all the queries that have been made to Cats
//type Query struct {
//        Id             int     // our internal id for this record
//        Image_Url      string  // url on the cat api
//        Source_Url     string  // the source (tumblr, imgur, redit, ...)
//}
