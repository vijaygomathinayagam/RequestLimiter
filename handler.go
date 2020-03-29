package main

import (
	"net"
	"net/http"
	"time"
	"strconv"
	"log"
)

func limitRequest(w http.ResponseWriter, req *http.Request) {
	// getting ip alone
	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		log.Printf("Error spliting request remote address, %q, %v\n", req.RemoteAddr, err)
	}
	userIP := net.ParseIP(ip)

	// getting access count from redis
	accessCount := getIPAccessCount(userIP.String())

	// checking ip is accessing within allowed limit
	if accessCount == ipAccessLimitCount {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// forward the request


	// increment the access count
	setIPAccessCount(userIP.String(), accessCount + 1)
}

func getIPAccessCount(ip string) int {
	// checking whether key exists
	existsVal, err := redisClient.Exists(ip).Result()
	if err != nil {
		log.Printf("Error while checking redis key: %s exists, err: %v\n", ip, err)
	}
	if existsVal == 0 {
		return 0
	}

	// reading the count
	var count int
	var countString string
	countString, err = redisClient.Get(ip).Result()
	if err != nil {
		log.Printf("Error while getting redis value for key: %s, err: %v\n", ip, err)
	}

	count, err = strconv.Atoi(countString)
	if err != nil {
		log.Printf("Error while converting redis value: %s to integer, err: %v\n", countString, err)
	}
	return count
}

func setIPAccessCount(ip string, count int) {
	err := redisClient.Set(ip, count, ipAccessLimitMinutes * time.Minute).Err()
	if err != nil {
		log.Printf("Error while setting redis value: %s, for key: %s, err: %v", count, ip, err)
	}
}