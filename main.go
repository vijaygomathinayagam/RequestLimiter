package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {
	http.HandleFunc("/", limitRequest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func limitRequest(w http.ResponseWriter, req *http.Request) {
	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		log.Printf("Error spliting request remote address, %q", req.RemoteAddr)
	}
	userIP := net.ParseIP(ip)
	fmt.Println(userIP)
}