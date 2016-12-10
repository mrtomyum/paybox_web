package controller

//import
import (
	"github.com/gorilla/websocket"
//"log"
	"net/http"
	"fmt"
	"github.com/paybox_terminal/model"
	"github.com/gin-gonic/gin"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Hub struct {
	Clients      map[*Client]bool
	//Broadcast    chan []byte
	Broadcast    chan model.Msg
	AddClient    chan *Client
	RemoveClient chan *Client
}

var Ghub = Hub{
	//Broadcast:    make(chan []byte),
	Broadcast:    make(chan model.Msg),
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
		case msg := <-hub.Broadcast:
			for conn := range hub.Clients {
				select {
				case conn.send <- msg:
				//default:
				//close(conn.send)
				//delete(hub.Clients, conn)
				}
			}
		}
	}
}

type Client struct {
	ws   *websocket.Conn
	//send chan []byte
	send chan model.Msg
}

func (c *Client) write() {
	defer func() {
		c.ws.Close()
	}()

	for {
		select {
		case msg, ok := <-c.send:
			if !ok {
				//  c.ws.WriteMessage(websocket.CloseMessage, []byte{})
				//  c.ws.WriteMessage(websocket.CloseMessage, []byte{})
				c.ws.WriteJSON(gin.H{"message" :"Connot to Send data" })
				return
			}

		//c.ws.WriteMessage(websocket.TextMessage, message)
			c.ws.WriteJSON(msg)
		}
	}
}

func (c *Client) read() {
	//	defer func() {
	//		Ghub.RemoveClient <- c
	//		c.ws.Close()
	//	}()
	msg := model.Msg{}
	for {
		//_, message, err := c.ws.ReadMessage()
		err := c.ws.ReadJSON(&msg)
		fmt.Println("command received : ", msg.Payload.Command)


		if err != nil {

			//Ghub.RemoveClient <- c
			fmt.Println("Ghub.RemoveClient working")
			fmt.Println("Formate Not working : ", msg)
			//c.ws.Close()
			c.ws.WriteJSON(gin.H{"Message":"invalid format received"})

			break
		}

		//todo : command  : onhand -> Get OnHandAmount and Bind data to payload & return to Client

		Ghub.Broadcast <- msg
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
		//send: make(chan []byte),
		send: make(chan model.Msg),
	}

	Ghub.AddClient <- client

	go client.write()
	go client.read()
}

//func homePage(res http.ResponseWriter, req *http.Request) {
//	http.ServeFile(res, req, "send.html")
//}
//func aa(res http.ResponseWriter, req *http.Request) {
//	http.ServeFile(res, req, "model.html")
//}
//func main() {
//	go Ghub.start()
//	http.HandleFunc("/ws", wsPage)
//	http.HandleFunc("/", homePage)
//	http.HandleFunc("/a", aa)
//	http.ListenAndServe(":8080", nil)
//}
