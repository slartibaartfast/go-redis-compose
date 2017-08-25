package main

import (
        "fmt"
        // s"time"
        "encoding/json"
        "strconv"
        "os"

        "github.com/garyburd/redigo/redis"
        // a"github.com/go-redis/redis"
)

//var currentPostId int
//var currentUserId int

func RedisConnect() redis.Conn {
  c, err := redis.Dial("tcp", os.Getenv("REDIS_URL"))
  //c, err := redis.Dial("tcp", "db:6379")
	HandleError(err)
	return c
}

// this uses import redis "gopkg.in/redis.v4"
//func NewClient() {
//  client := redis.NewClient(&redis.Options{
//      Addr:     os.Getenv("REDIS_URL"),
//      Password: "", // no password set
//      DB:       0,  // use default DB
//  })
//}

// Add seed data
func init() {
	CreateCat(Cat{
		Id: 1,
		Url: "http://24.media.tumblr.com/tumblr_lyhac7a3lJ1r4zr2vo1_r1_500.gif",
		Source_Url: "http://thecatapi.com/?id=4ov",
	})

  CreateCat(Cat{
		Id: 2,
		Url: "http://25.media.tumblr.com/tumblr_m20fahuHFk1qzex9io1_1280.jpg",
		Source_Url: "http://thecatapi.com/?id=q6",
	})
}


func CreateCat(cat Cat) {

  currentCatId := cat.Id + 1
	//currentUserId += 1

	cat.Id = currentCatId
	//cat.Image_Url = currentUserId
	//cat.Source_Url = time.Now()

  // connect to redis
	c := RedisConnect()
	defer c.Close()

  // set up our data as a blob
	b, err := json.Marshal(cat)
	HandleError(err)

	// Save JSON blob to Redis
	reply, err := c.Do("SET", "cat:" + strconv.Itoa(cat.Id), b)
	HandleError(err)

  // get the redis(?) reply
	fmt.Println("GET ", reply)
}


// Return all the records from Cat
func FindAll() Cats {

	var cats Cats

	c := RedisConnect()
	defer c.Close()

	keys, err := c.Do("KEYS", "cat:*")
	HandleError(err)

	for _, k := range keys.([]interface{}) {

		var cat Cat

		reply, err := c.Do("GET", k.([]byte))
		HandleError(err)

		if err := json.Unmarshal(reply.([]byte), &cat); err != nil {
			panic(err)
		}
		cats = append(cats, cat)
	}
	return cats
}

//func FindPost(id int) Post {

//	var post Post

//	c := RedisConnect()
//	defer c.Close()

//	reply, err := c.Do("GET", "post:" + strconv.Itoa(id))
//	HandleError(err)
//
//	fmt.Println("GET OK")

//	if err = json.Unmarshal(reply.([]byte), &post); err != nil {
//		panic(err)
//	}
//	return post
//}
