package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"

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
		case client := <-r.join:
			r.clients[client] = true

		case client := <-r.leave:
			delete(r.clients, client)
			close(client.receive)
		case message := <-r.forward:
			for client := range r.clients {
				client.receive <- message
			}
		}
	}
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}

var rooms = make(map[string]*room)

var mu sync.Mutex

func getRoom(name string) *room {
	mu.Lock()
	defer mu.Unlock()

	if r, ok := rooms[name]; ok {
		return r
	}

	r := newRoom()
	rooms[name] = r
	go r.run()
	return r
}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	roomName := req.URL.Query().Get("room")
	if roomName == "" {
		http.Error(w, "room name is required", http.StatusBadRequest)
		return
	}

	realRoom := getRoom(roomName)

	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}

	client := &client{
		socket:  socket,
		receive: make(chan []byte, messageBufferSize),
		room:    realRoom,
		name:    fmt.Sprintf("Guest_%d", rand.Intn(100000)),
	}
	realRoom.join <- client
	defer func() { realRoom.leave <- client }()
	go client.write()
	client.read()
}
