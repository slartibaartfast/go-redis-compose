# go-redis-compose
Docker compose for Redis and Golang

This is a practice Go project that receives xml from an external api and transforms it to JSON.  Data is stored in Redis.  Redis lives in one container, and the Go api lives in another.

The containers are based on Alpine Linux 3.6.  The container for the api has a few extra packages added to it, namely:
- github.com/julienschmidt/httprouter
- github.com/garyburd/redigo/redis

These projects are also in the /vendor subdirectory, govendor style.


===========


## Running locally

Build and run using Swarm and Compose:

If you don't already have one, get a key for the venerable cat api from
[http://thecatapi.com/api-key-registration.html](http://thecatapi.com/api-key-registration.html)

Create a file, go-redis-compose/api-key.env, and save your cat api key in it.  This gets used later as a Docker Secrets file.  While it is a little silly to treat a public api key as a secret, it seems like a safe value to use when first integrating secrets into a project.

Open docker-compose.yml and edit the api Volumes so that your absolute path gets mapped.

>Volumes:
>
>  \- </your/absolute/path/>/go-redis-compose: /code/go/src/go-redis-compose

Start a swarm engine:

  `$ docker swarm init`


Deploy the stack with the name cats:

  `$ docker stack deploy --compose-file=docker-compose.yml cats`


To see the status of service creation, run:

  `$ docker service ls`


Once pulls are done and everything is built, check [http://0.0.0.0:8080/cat](http://0.0.0.0:8080/cat) to view a json string related to an image of a cat, retrieved from thecatapi.com

Check [http://0.0.0.0:8080/history](http://0.0.0.0:8080/history) to view the cumulative set of calls made to /cat during the session.


To bring the stack down:

  `$ docker stack rm cats`
