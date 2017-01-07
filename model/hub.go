package model

import "fmt"

type Hub struct {
	Clients   []*Client
	AddClient chan *Client
	DelClient chan *Client
	Send      chan *Client
}

func (hub *Hub) Start() {
	for {
		select {
		case c := <-hub.AddClient:
			fmt.Println("hub.Start.AddClient Working....")
			hub.Clients = append(hub.Clients, c)
		case c := <-hub.DelClient:
			fmt.Println("hub.Start.RemoveClient Working....")
		//if _, ok := hub.Clients[c]; ok {
			for key, activeClient := range hub.Clients {
				if c == activeClient {
					hub.Clients = append(hub.Clients[:key], hub.Clients[key + 1:]...) // delete slice of client
					close(c.Send)
				}
			}

		//case m := <-hub.Broadcast:
		//	for c := range hub.Clients {
		//		select {
		//		case c.Send <- m:
		//			fmt.Println("Broadcast msg to View.Send")
		//		//default:
		//		//close(conn.send)
		//		//delete(hub.Clients, conn)
		//		}
		//	}

		}
	}
}
