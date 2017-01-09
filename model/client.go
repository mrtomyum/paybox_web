package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
)

type Client struct {
	Ws   *websocket.Conn
	Send chan *Message
	Name string
	Msg  *Message
}

func (c *Client) Read() {
	defer func() {
		c.Ws.Close()
	}()
	m := &Message{}
	for {
		err := c.Ws.ReadJSON(&m)
		if err != nil {
			log.Println("Error ReadJSON():", err)
			return
		}
		c.Msg = m
		switch c.Name {
		case "web":
			fmt.Println("Message from web")
			switch m.Command {
			case "onhand":
				H.Onhand(c)
			case "cancel":
				H.Cancel(c)
			case "billing":
				H.Billing(c)
			}
		case "dev":
			fmt.Println("dev")
			switch m.Device {
			case "coin_hopper":
			case "coin_acc":
			case "bill_acc":
			case "printer":
			}
		//return
		default:
			fmt.Println("Case default: Message==>", m)
			m.Type = "response"
			m.Data = "Hello"
			c.Send <- m
		}
	}
}

func (c *Client) Write() {
	defer func() {
		c.Ws.Close()
	}()
	for {
		select {
		case m, ok := <-c.Send:
			if !ok {
				c.Ws.WriteJSON(gin.H{"message": "Cannot send data"})
				return
			}
			fmt.Println("Client.Write():", m)
			c.Ws.WriteJSON(m)
		}
	}
}
