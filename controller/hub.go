package controller

//import
import (
	"github.com/gorilla/websocket"
//"log"
	"net/http"
	"fmt"
	"github.com/mrtomyum/paybox_terminal/model"
	"github.com/gin-gonic/gin"
//"strconv"
//	"strconv"
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


var onHand = model.OnHand{
	OnhandAmount : 0,
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
		switch
		{
		case msg.Payload.Command == "cancel":
			// Reset Onhand
			// todo: must be send return money to Hardware
			onHand.OnhandAmount = 0

			res := model.Msg{}
			res.Payload.Command = "cancel"
			res.Payload.Data = "Cancel - Successful"
			res.Payload.Result = true
			res.Device = "Host"
			res.Payload.Type = "response"
			Ghub.Broadcast <- res


		case msg.Payload.Command == "billing":
			// todo: save into databse sqlite
			res := model.Msg{}
			res.Payload.Command = "billing"
			res.Payload.Data = "Docno : xxxxxx sucessful"
			res.Payload.Result = true
			res.Device = "Host"
			res.Payload.Type = "response"
			Ghub.Broadcast <- res

		// for Client - UI/UX call check current onhand amount
		case msg.Payload.Command == "onhand" && msg.Payload.Type == "request":
			res := model.Msg{}
			res.Payload.Command = "onhand"
			res.Payload.Data = onHand.OnhandAmount
			res.Payload.Result = true
			res.Device = "Host"
			res.Payload.Type = "response"
			Ghub.Broadcast <- res

		// for push  totalAmount Update from hardware event and sum new onhand amount and send to UI
		case msg.Payload.Command == "onhand" && msg.Payload.Type == "event" :
			res := model.Msg{}
			res.Payload.Command = "onhand"
			fmt.Println("onhand_event_starting....")
			//ปรับยอด Onhand ตามเงินที่เข้ามา

			// bind interface{} to i variable
			i := msg.Payload.Data


			//check type of interface{}
			//int_amount, _ := amount.(int)

			switch  i.(type) {
			case float64:
				fmt.Println("amnount type : float64")
			case float32:
				fmt.Println("amnount type : float32")
			case int64:
				fmt.Println("amnount type : int64")

			}


			// convert interface{} to int
			// example onhand_event
			var iAreaId int = int(i.(float64))

			// Update Current OnHand TOTAL
			onHand.OnhandAmount = onHand.OnhandAmount + iAreaId
			fmt.Println("Current Totalamount : ", onHand.OnhandAmount)

			res.Payload.Data = onHand.OnhandAmount
			res.Payload.Result = true
			res.Device = "Host"
			res.Payload.Type = "response"
			Ghub.Broadcast <- res
		default :
			Ghub.Broadcast <- msg
		}

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

