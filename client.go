package main

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

type client struct {
	socket *websocket.Conn

	receive chan []byte

	room *room

	name string
}

func (c *client) read() {

	defer c.socket.Close()

	for {
		_, message, err := c.socket.ReadMessage()
		if err != nil {
			return
		}
		outgoing := map[string]string{
			"name":    c.name,
			"message": string(message),
		}
		jsMessage, err := json.Marshal(outgoing)

		if err != nil {
			fmt.Println("Encoding error", err)
			continue
		}
		c.room.forward <- jsMessage
	}
}

func (c *client) write() {
	defer c.socket.Close()
	for message := range c.receive {
		err := c.socket.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			return
		}
	}
}
