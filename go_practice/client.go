package main

import "github.com/gorilla/websocket"

type client struct {
	to     chan []byte
	room   *room
	socket *websocket.Conn
}

func (c *client) ReadMessage() {
	defer c.socket.Close()
	for {
		_, msg, err := c.socket.ReadMessage()		
		if err != nil {
			return
		}

		c.room.forward <- msg
	}
}

func (c *client) SendMessage() {
	defer c.socket.Close()
	
	for msg := range c.to {
		err := c.socket.WriteMessage(websocket.TextMessage, msg) 
		if err != nil {
			break
		}
	}
}
