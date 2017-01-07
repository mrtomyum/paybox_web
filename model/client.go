package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
)

type Client struct {
	Conn *websocket.Conn
	Send chan Msg
	Name string
}

func (c *Client) Write() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		select {
		case msg, ok := <-c.Send:
			if !ok {
				c.Conn.WriteJSON(gin.H{"message" :"Connot to Send data" })
				return
			}
			fmt.Println("View.Write():", msg)
		//c.Conn.WriteJSON(msg)
		}
	}
}

func (c *Client) Read() {
	m := Msg{}
	for {
		err := c.Conn.ReadJSON(&m)
		if err != nil {
			log.Println("Client Read JSON Error:", m)
			c.Conn.WriteJSON(gin.H{"message": "Read JSON Error: "})
			break
		}
		fmt.Println("Received command from Client: ", m)

		switch m.Payload.Command {
		// for Client - UI/UX call check current onhand amount
		case "onhand":
			host.GetOnHand(c, m)

		// Cancel order command
		// อย่าลืม reset ยอดเงินในตู้เป็น 0 (hardware ต้องสั่งคืนเงิน)
		case "cancel":
			host.Cancel(c, m)

		// billing command ใช้สำหรับให้ Client เรียกบันทึกเข้ามาที่ Websocket
		case "billing":
			host.Billing(c, m)

		//default:
		//	MyHub.Broadcast <- msg
		}

	}
}
