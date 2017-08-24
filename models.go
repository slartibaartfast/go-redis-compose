package main


//TODO: you might only need the cat struct.  Just add all queries to the external
// cat api to it, use it to return them one at the /cat endpoints
// FindAll to give them the history endpoint

// cats structure
type Cat struct {
        Id           int     `json:"id"`          // our internal id
        Url          string  `json:"url"`         // the original url (tumblr, imgur, redit, ...)
        Source_Url   string  `json:"source_url"`  // url of the image on the cat api
}


// history structure, all the queries that have been made to Cats
//type Query struct {
//        Id             int     // our internal id for this record
//        Image_Url      string  // url on the cat api
//        Source_Url     string  // the source (tumblr, imgur, redit, ...)
//}

type Cats []Cat
