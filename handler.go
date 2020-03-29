package main

import (
	"net"
	"net/http"
	"time"
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
	var count int
	err := redisClient.Get(ip).Scan(count)
	if err != nil {
		log.Printf("Error while scanning redis value for key: %s, err: %v\n", ip, err)
	}
	return count
}

func setIPAccessCount(ip string, count int) {
	err := redisClient.Set(ip, count, ipAccessLimitMinutes * time.Minute).Err()
	if err != nil {
		log.Printf("Error while setting redis value: %s, for key: %s, err: %v", count, ip, err)
	}
}