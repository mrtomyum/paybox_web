package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Client struct {
	Conn *websocket.Conn
	Send chan Msg
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
			fmt.Println(msg)
			c.Conn.WriteJSON(msg)
		}
	}
}

func (c *Client) Read() {
	msg := Msg{}
	for {
		err := c.Conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println("Client Read JSON Error:", msg)
			c.Conn.WriteJSON(gin.H{"message": "Read JSON Error: "})
			break
		}
		fmt.Println("Received command from Client: ", msg)

		switch msg.Payload.Command {
		// for Client - UI/UX call check current onhand amount
		case "onhand":
			host.GetOnHand(c, msg)

		// Cancel order command
		// อย่าลืม reset ยอดเงินในตู้เป็น 0 (hardware ต้องสั่งคืนเงิน)
		case "cancel":
			host.Cancel(c, msg)

		// billing command ใช้สำหรับให้ Client เรียกบันทึกเข้ามาที่ Websocket
		case "billing":
			host.Billing(c, msg)

		default:
			Ghub.Broadcast <- msg
		}

	}
}
