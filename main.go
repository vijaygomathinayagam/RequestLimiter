package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"net/http"
)

const (
	redisHost            = "redis"
	redisPort            = "6379"
	appListenPort        = "8080"
	ipAccessLimitMinutes = 1
	ipAccessLimitCount   = 100
)

var (
	redisClient *redis.Client
)

func main() {
	initRedis()
	http.HandleFunc("/", limitRequest)
	fmt.Printf("Application is listening on port: %s\n", appListenPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", appListenPort), nil))
}

func initRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", redisHost, redisPort),
	})
}
