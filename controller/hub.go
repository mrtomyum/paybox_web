package controller

import (
	"github.com/gorilla/websocket"
	"github.com/mrtomyum/paybox_terminal/model"
)

type Hub struct {
	clients      map[*Client]bool
	addClient    chan *Client
	removeClient chan *Client
}

var hub = Hub{
	clients: make(map[*Client]bool),
	addClient: make(chan *Client),
	removeClient: make(chan *Client),
}

func (hub *Hub) start() {
	for {
		select {
		case conn := <-hub.addClient:
			hub.clients[conn] = true
		case conn := <-hub.removeClient:
			if _, ok := hub.clients[conn]; ok {
				delete(hub.clients, conn)
				close(conn.send)
			}
		}
	}
}

type Client struct {
	ws   *websocket.Conn
	send chan []byte
}

func (c *Client) write() {
	defer func() {
		c.ws.Close()
	}()

	for {
		select {
		case msg, ok := <-c.send:
			if !ok {
				c.ws.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
		//c.ws.WriteMessage(websocket.TextMessage, msg)
			c.ws.WriteJSON(msg)
		}
	}
}

func (c *Client) read() {
	defer func() {
		hub.removeClient <- c
		c.ws.Close()
	}()

	msg := model.Msg{}

	for {
		err := c.ws.ReadJSON(&msg)
		if err != nil {
			hub.removeClient <- c
			c.ws.Close()
			break
		}
		// Do something with msg.
	}
}
