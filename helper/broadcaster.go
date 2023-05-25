package helper

import "log"

func NewBroadcaster() *Broadcaster {
	return &Broadcaster{
		Broadcast:     make(chan []byte),
		AddChannel:    make(chan chan []byte),
		RemoveChannel: make(chan chan []byte),
		channels:      make(map[chan []byte]int32),
	}
}

type Broadcaster struct {
	// Channel for incoming messages to broadcast
	Broadcast chan []byte

	// Channel to add a new channel
	AddChannel chan chan []byte

	// Channel to remove a channel
	RemoveChannel chan chan []byte

	// Channels
	channels map[chan []byte]int32
}

func (b *Broadcaster) Listen() {
	var seq int32 = 0

	for {
		select {
		case message := <-b.Broadcast:
			// broadcast message to all channels
			for channel := range b.channels {
				channel <- message
			}
		case channel := <-b.AddChannel:
			// Add a new listener
			b.channels[channel] = seq
			seq++
			log.Println("New channel added")
		case channel := <-b.RemoveChannel:
			// Close a channel
			delete(b.channels, channel)
			log.Println("Deleted channel")
		}
	}
}
