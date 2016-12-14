package model

var Ghub = Hub{
	//Broadcast:    make(chan []byte),
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
			hub.Clients[conn] = true
		case conn := <-hub.RemoveClient:
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
