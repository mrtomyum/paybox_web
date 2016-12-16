package model
import "fmt"

var Ghub = Hub{
	Broadcast:    make(chan Msg),
	AddClient:    make(chan *Client),
	RemoveClient: make(chan *Client),
	Clients:      make(map[*Client]bool),
}

type Hub struct {
	Clients      map[*Client]bool
	Broadcast    chan Msg
	AddClient    chan *Client
	RemoveClient chan *Client
}

func (hub *Hub) Start() {
	for {
		select {
		case conn := <-hub.AddClient:
			fmt.Println("hub.Start.AddClient Working....")
			hub.Clients[conn] = true
		case conn := <-hub.RemoveClient:
			fmt.Println("hub.Start.RemoveClient Working....")
			if _, ok := hub.Clients[conn]; ok {
				delete(hub.Clients, conn)
				close(conn.Send)
			}
		case msg := <-hub.Broadcast:
			for conn := range hub.Clients {
				select {
				case conn.Send <- msg:
				//default:
				//close(conn.send)
				//delete(hub.Clients, conn)
				}
			}
		}
	}
}
