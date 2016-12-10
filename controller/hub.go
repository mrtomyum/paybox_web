package controller

//import
import (
	"github.com/gorilla/websocket"
//"log"
	"net/http"
	"fmt"
)

var upgrader = websocket.Upgrader{}

type Hub struct {
	Clients      map[*Client]bool
	Broadcast    chan []byte
	AddClient    chan *Client
	RemoveClient chan *Client
}

var Ghub = Hub{
	Broadcast:    make(chan []byte),
	AddClient:    make(chan *Client),
	RemoveClient: make(chan *Client),
	Clients:      make(map[*Client]bool),
}

func (hub *Hub) Start() {
	for {
		select {
		case conn := <-hub.AddClient:
			hub.Clients[conn] = true
		case conn := <-hub.RemoveClient:
			if _, ok := hub.Clients[conn]; ok {
				delete(hub.Clients, conn)
				close(conn.send)
			}
		case message := <-hub.Broadcast:
			for conn := range hub.Clients {
				select {
				case conn.send <- message:
				default:
					close(conn.send)
					delete(hub.Clients, conn)
				}
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
		case message, ok := <-c.send:
			if !ok {
				c.ws.WriteMessage(websocket.CloseMessage, []byte{})

				return
			}

			c.ws.WriteMessage(websocket.TextMessage, message)
		}
	}
}

func (c *Client) read() {
	defer func() {
		Ghub.RemoveClient <- c
		c.ws.Close()
	}()

	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			Ghub.RemoveClient <- c
			c.ws.Close()
			break
		}

		Ghub.Broadcast <- message
	}
}

func wsPage(res http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(res, req, nil)
	fmt.Println("ws : wsPage start")
	if err != nil {
		http.NotFound(res, req)
		return
	}

	client := &Client{
		ws:   conn,
		send: make(chan []byte),
	}

	Ghub.AddClient <- client

	go client.write()
	go client.read()
}

func homePage(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "send.html")
}
func aa(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "model.html")
}
//func main() {
//	go Ghub.start()
//	http.HandleFunc("/ws", wsPage)
//	http.HandleFunc("/", homePage)
//	http.HandleFunc("/a", aa)
//	http.ListenAndServe(":8080", nil)
//}
