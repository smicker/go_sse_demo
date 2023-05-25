package helper

import "log"

func NewBroadcaster() *Broadcaster {
	return &Broadcaster{
		Broadcast:       make(chan []byte),
		NewConnection:   make(chan chan []byte),
		CloseConnection: make(chan chan []byte),
		connections:     make(map[chan []byte]int32),
	}
}

type Broadcaster struct {
	// Channel for incoming messages to broadcast
	Broadcast chan []byte

	// Channel to accept new connections
	NewConnection chan chan []byte

	// Channel to close connections
	CloseConnection chan chan []byte

	// Connections
	connections map[chan []byte]int32
}

func (b *Broadcaster) Listen() {
	var seq int32 = 0

	for {
		select {
		case message := <-b.Broadcast:
			// broadcast message to all connections
			for connection := range b.connections {
				connection <- message
			}
		case connection := <-b.NewConnection:
			// Add a new connection
			b.connections[connection] = seq
			seq++
			log.Println("New connection added")
		case connection := <-b.CloseConnection:
			// Close a connection
			delete(b.connections, connection)
			log.Println("Deleted Connection")
		}
	}
}
