package main

import (
	"encoding/xml"
)

// cat structure, for storing json related to images of cats
type Cat struct {
	Id         string `json:"id"`         // the id from the cat api (used to be our internal int id)
	Url        string `json:"url"`        // the original url (tumblr, imgur, redit, ...)
	Source_Url string `json:"source_url"` // url of the image on the cat api
}

type Image struct {
	Image Cat `json:"image"` // let data from a cat struct be an Image
}

type Images struct {
	Images []Cat `json:"images"` //let a dictionary of cat data be Images
}

// structure to match the cat api xml output so that it may be mapped to
type xmlResponse struct {
	Response xml.Name   `xml:"response"`
	Images   []xmlImage `xml:"data>images>image"`
}

// cat image data xml structure per thecatapi.com schema
type xmlImage struct {
	Image      xml.Name `xml:"image"`
	Url        string   `xml:"url"`
	Id         string   `xml:"id"`
	Source_Url string   `xml:"source_url"`
}

// slice of Cat
type Cats []Cat
