/*
An http server that sends messages to connected clients through Server-Sent Events (SSE).
	* Only string can be sent.
	* Supports a maximum of 6 SSE channels per browser session.
*/

package sse

import (
	"broadcast_demo/helper"
	"fmt"
	"net/http"
)

func NewServer() *Server {
	broadcaster := helper.NewBroadcaster()
	server := &Server{
		broadcaster,
	}
	go broadcaster.Listen()
	return server
}

func setHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

type Server struct {
	broadcaster *helper.Broadcaster
}

func (server *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, ok := w.(http.Flusher)

	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	setHeaders(w)
	clientConnection := make(chan []byte)
	server.broadcaster.NewConnection <- clientConnection
	defer func() {
		server.broadcaster.CloseConnection <- clientConnection
	}()

	isRequestClosed := r.Context().Done()

	go func() {
		<-isRequestClosed
		server.broadcaster.CloseConnection <- clientConnection
	}()

	for {
		// This will send any message that comes in on the clientConnection channel to the SSE.
		message := <-clientConnection
		sendSseMessage(w, message)
	}
}

// Send message to SSE.
func sendSseMessage(w http.ResponseWriter, message []byte) {
	flusher, _ := w.(http.Flusher)
	// For Server-Sent Events (SSE) you always have to send a utf-8 encoded string and it
	// must always start with "data: ". You cannot send something like images.
	fmt.Fprintf(w, "data: %s\n\n", message)
	flusher.Flush()
}

// Broadcast a message to all listening channels
func (server *Server) Broadcast(message string) {
	server.broadcaster.Broadcast <- []byte(message)
}
