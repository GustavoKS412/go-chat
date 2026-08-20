package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/gorilla/websocket"
)

type room struct {
	clients map[*client]bool

	join chan *client

	leave chan *client

	forward chan []byte
}

func newRoom() *room {
	return &room{
		forward: make(chan []byte),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
	}
}

func (r *room) run() {
	for {
		select {
		case c := <-r.join:
			r.clients[c] = true

		case c := <-r.leave:
			delete(r.clients, c)
			close(c.receive)
		case msg := <-r.forward:
			for c := range r.clients {
				c.receive <- msg
			}
		}
	}
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}

	c := &client{
		socket:  socket,
		receive: make(chan []byte, messageBufferSize),
		room:    r,
		name:    fmt.Sprintf("Guest_%d", rand.Intn(100000)),
	}
	r.join <- c
	defer func() { r.leave <- c }()
	go c.write()
	c.read()
}
