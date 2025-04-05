package main

import (
	"fmt"
	"github.com/mmarci96/go-ws-traffic-controller/go-reverse-proxy/internal/server"
	"log"
)

func main() {
	fmt.Println("Reverse proxy starting...")
	addr, err := server.Run()
	if err != nil {
		log.Fatalf("could not start the server: %v", err)
	}
	fmt.Printf("Proxy server running on:http://%s\n", addr)
	select {}
}
