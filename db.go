package main

import (
        "fmt"
        // s"time"
        "encoding/json"
        // s"strconv"
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
  // connect to redis
	c := RedisConnect()
	defer c.Close()

  // purge any existing data
  reply, err := c.Do("FLUSHALL")
  HandleError(err)
  fmt.Println("FLUSHALL ", reply)


//	CreateCat(Cat{
//		Id: "asdf",
//		Url: "http://24.media.tumblr.com/tumblr_lyhac7a3lJ1r4zr2vo1_r1_500.gif",
//		Source_Url: "http://thecatapi.com/?id=4ov",
//	})

//  CreateCat(Cat{
//		Id: "rtwe",
//		Url: "http://25.media.tumblr.com/tumblr_m20fahuHFk1qzex9io1_1280.jpg",
//		Source_Url: "http://thecatapi.com/?id=q6",
//	})
}


func CreateCat(cat Cat) {

  //currentCatId := cat.Id
	//currentUserId += 1

	//cat.Id = currentCatId

  fmt.Println(cat.Id)
  //currentCatId, err := strconv.Atoi(cat.Id)
  //if err != nil {
  // panic(err)
  //}
  //fmt.Println(currentCatId)

  // connect to redis
	c := RedisConnect()
	defer c.Close()

  // set up our data as a blob
	b, err := json.Marshal(cat)
	HandleError(err)

	// Save JSON blob to Redis
	reply, err := c.Do("SET", "cat:" + cat.Id, b)
  fmt.Println("SET", "cat:" + cat.Id, b)
  HandleError(err)

  // get the redis(?) reply
	fmt.Println("GET ", reply)
}


// Return all the records from Cat
//func FindAll() Cats {
func FindAll() Images {

  //var cats Images
  var tmp Cats
  //var cat Cat

	c := RedisConnect()
	defer c.Close()

	keys, err := c.Do("KEYS", "cat:*")
	HandleError(err)

  //if err := json.Unmarshal(keys.([]byte), &cats); err != nil {
  //  panic(err)
  //}

  //map[cats.Images][]cat{}
	for _, k := range keys.([]interface{}) {

		var cat Cat

    fmt.Println("k ", k)
    //fmt.Println("k ", string(k))

		reply, err := c.Do("GET", k.([]byte))
		HandleError(err)
    fmt.Println("reply ", reply)

    if err := json.Unmarshal(reply.([]byte), &cat); err != nil {
			panic(err)
		}
    //if err := json.Unmarshal(reply.([]byte), &cats.Images.Cat); err != nil {
    //  panic(err)
    //}
		//cats = append(cats, cat)
    tmp = append(tmp, cat)
    //cats = append(cats.Images, cat)
    //cats = cats.Images.append(cat)
	}
	//return cats
  return Images{tmp}
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
