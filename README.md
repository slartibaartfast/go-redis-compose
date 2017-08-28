# go-redis-compose
Docker compose for redis and go

This is a practice Go project that receives xml from an external api and transforms it to JSON.  Data is stored in Redis.  Redis lives in one container, and the api lives in another.

The containers are based on Alpine Linux.  The container for the api has a few extra packages added to it, namely:
- github.com/julienschmidt/httprouter
- github.com/garyburd/redigo/redis

These projects are also in the /vendor subdirectory, govendor style.


===========


## Running locally

Build and run using Docker Compose:

If you don't already have one, get a key for the venerable cat api from
[http://thecatapi.com/api-key-registration.html](http://thecatapi.com/api-key-registration.html)


	$ docker-compose up


Check [http://0.0.0.0:8080/cat](http://0.0.0.0:8080/cat) to view a json string related to an image of a cat, retrieved from thecatapi.com

Check [http://0.0.0.0:8080/history](http://0.0.0.0:8080/history) to view the cumulative set of calls made to /cat during the session.
