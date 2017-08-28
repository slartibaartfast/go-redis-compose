package main

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"os"
)

// connect to our redis database
func RedisConnect() redis.Conn {
	c, err := redis.Dial("tcp", os.Getenv("REDIS_URL"))
	HandleError(err)
	return c
}

// Purge old data and add seed data for testing
func init() {
	// connect to redis
	c := RedisConnect()
	defer c.Close()

	// purge any existing data
	reply, err := c.Do("FLUSHALL")
	HandleError(err)
	fmt.Println("FLUSHALL ", reply)

	// add some data for testing
	//	CreateCat(Cat{
	//		Id: "4ov",
	//		Url: "http://24.media.tumblr.com/tumblr_lyhac7a3lJ1r4zr2vo1_r1_500.gif",
	//		Source_Url: "http://thecatapi.com/?id=4ov",
	//	})

	//  CreateCat(Cat{
	//		Id: "q6",
	//		Url: "http://25.media.tumblr.com/tumblr_m20fahuHFk1qzex9io1_1280.jpg",
	//		Source_Url: "http://thecatapi.com/?id=q6",
	//	})
}

// save data in Cat structure to redis
func CreateCat(cat Cat) {
	// connect to redis
	c := RedisConnect()
	defer c.Close()

	// set up our data as a blob of json strings
	b, err := json.Marshal(cat)
	HandleError(err)

	// save json blob to redis
	reply, err := c.Do("SET", "cat:"+cat.Id, b)
	fmt.Println("SET", "cat:"+cat.Id, b)
	HandleError(err)

	// get the redis reply
	fmt.Println("GET ", reply)
}

// return all the cats in an Images structure
func FindAll() Images {
	// let a slice of Cat, the struct Cats, hold values
	var tmp Cats

	// connect to the database
	c := RedisConnect()
	defer c.Close()

	// get all the data with a key like cat:% ... every cat key value pair
	// TODO: use SCAN instead.  see https://redis.io/commands/keys
	keys, err := c.Do("KEYS", "cat:*")
	HandleError(err)

	// loop though the returned data and load it into an Images structure
	for _, k := range keys.([]interface{}) {

		// for storing Cat specific data
		var cat Cat

		// get the value for this key
		reply, err := c.Do("GET", k.([]byte))
		HandleError(err)

		// load the data into a Cat structure
		if err := json.Unmarshal(reply.([]byte), &cat); err != nil {
			panic(err)
		}

		// add the cat data to the slice
		tmp = append(tmp, cat)
	}
	// add the slice to the Images struct and return all the data as Images
	return Images{tmp}
}
