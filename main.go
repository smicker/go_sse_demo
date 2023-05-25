package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"broadcast_demo/sse"
)

func main() {
	fmt.Println("Starting main")

	sseServer := sse.NewServer()

	startPinging(sseServer)

	log.Println("Starting")
	log.Println("Error: ", http.ListenAndServe("localhost:8080", sseServer))
}

// Broadcast a message every 3 seconds
func startPinging(server *sse.Server) {
	ticker := time.NewTicker((3 * time.Second))

	quit := make(chan struct{}) // The sender to quit channel is not implemented in this demo (out of scope)!!

	go func() {
		for {
			select {
			case <-ticker.C:
				randomNumber := strconv.Itoa(rand.Intn(10))
				log.Printf("Broadcasting %s", randomNumber)
				server.Broadcast(randomNumber)
			case <-quit:
				log.Println("Stopping pinger!!")
				ticker.Stop()
				return
			}
		}
	}()
}
