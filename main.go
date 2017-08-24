package main

import (
    // "fmt"
    // "os"
    // "github.com/go-redis/redis"

    // redis "gopkg.in/redis.v4"

    "log"
  	"net/http"
)

//func main() {
//    client := redis.NewClient(&redis.Options{
//        Addr:     os.Getenv("REDIS_URL"),
//        Password: "", // no password set
//        DB:       0,  // use default DB
//    })

//    pong, err := client.Ping().Result()
//    fmt.Println(pong, err)
//}

func main() {
    	  router := NewRouter()

	      log.Fatal(http.ListenAndServe(":8080", router))
}


func HandleError(err error) {
        if err != nil {
                panic(err)
        }
}
