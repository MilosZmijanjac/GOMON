package main

import (
	"log"
	"sync"

	"go-micro.dev/v5/broker"
)

type Hub struct {
	sync.RWMutex

	clients map[string]*Client

	broadcast  chan *Message
	register   chan *Client
	unregister chan *Client
	direct     chan *Message
	publish    chan *Message
}

func NewHub() *Hub {
	return &Hub{
		clients:    map[string]*Client{},
		broadcast:  make(chan *Message),
		publish:    make(chan *Message),
		direct:     make(chan *Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) run() {
	if err := broker.Connect(); err != nil {
		log.Fatalf("Broker connect error: %v", err)
	}
	for {
		select {
		case client := <-h.register:
			h.Lock()
			h.clients[client.id] = client
			h.Unlock()

			log.Printf("client registered %s", client.id)
		case client := <-h.unregister:
			h.Lock()
			if _, ok := h.clients[client.id]; ok {
				close(client.send)
				log.Printf("client unregistered %s", client.id)
				delete(h.clients, client.id)
			}
			h.Unlock()
		case msg := <-h.direct:
			h.RLock()
			log.Printf("client direct %s", msg.ClientID)
			h.clients[msg.ClientID].send <- []byte(msg.Text)
			h.RUnlock()
		case msg := <-h.publish:
			h.RLock()
			log.Printf("client publish %s", msg.ClientID)
			pbs := &broker.Message{
				Header: map[string]string{"id": "1"},
				Body:   []byte(msg.Text),
			}
			if err := broker.Publish("device.data", pbs); err != nil {
				log.Printf("Error publishing: %v", err)
			}
			h.RUnlock()
		case msg := <-h.broadcast:
			h.RLock()
			log.Printf("client broadcast %s", msg.ClientID)
			for client := range h.clients {
				select {
				case h.clients[client].send <- []byte(msg.Text):
				default:
					close(h.clients[client].send)
					delete(h.clients, client)
				}
			}
			h.RUnlock()

		}

	}
}
