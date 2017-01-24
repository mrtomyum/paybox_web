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
			log.Println("Connection closed:", err)
			break
		}
		c.Msg = m
		switch {
		case c.Name == "web":
			c.WebEvent()
		case c.Name == "dev":
			c.DevEvent()
		default:
			fmt.Println("Error: Case default: Message==>", m)
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
			fmt.Println("Client.Write():", c.Name, m)
			c.Ws.WriteJSON(m)
		}
	}
}

// WebEvent แยกเส้นทาง Message Request จาก Web Frontend โดยแยกตาม Command ต่างๆ
func (c *Client) WebEvent() {
	// ปกติแล้ว  Web จะไม่สั่งการ Device ตรงๆ
	// จะสั่งผ่าน Host ให้ Host ทำงานระดับล่างแทน
	// แต่ตรงนี้มีไว้สำหรับการ Debug ผ่าน Web GUI
	fmt.Println("Request message from Web")
	switch c.Msg.Command {
	case "onhand":
		PM.OnHand(c)
	case "cancel":
		PM.Cancel(c)
	default:
		log.Println("Client.WebEvent(): default: Unknown Command for web client=>", c.Msg.Command)
	}
}

// DevEvent เป็นการแยกเส้นทาง Message จาก Device Event และ Response
// โดย Function นี้จะแยก message ตาม Device ก่อน แล้วจึงแยกเส้นทางตาม Command
func (c *Client) DevEvent() {
	fmt.Println("method DevEvent() Event message from Dev:", c.Msg)
	switch c.Msg.Device {
	case "coin_hopper":
		CH.Event(c)
	case "coin_acc":
		CA.Event(c)
	case "bill_acc":
		BA.Event(c)
	case "printer":
		P.Event(c)
	case "main_board":
		M.Event(c)
	default:
		log.Println("event cannot find function/message=", c.Msg)
	}
}
