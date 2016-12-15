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

var machine = Machine{
	Id:     "1",
	OnHand: 0,
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
			c.Conn.WriteJSON(msg)
		}
	}
}

//func (c *Client) Read() {
//	msg := Msg{}
//	for {
//		// รับ msg ในรูป JSON เข้ามาจากไหนยังไม่รู้
//		err := c.Conn.ReadJSON(&msg)
//		fmt.Println("command received : ", msg.Payload.Command)
//		if err != nil {
//			fmt.Println("Ghub.RemoveClient working")
//			fmt.Println("Format Not working : ", msg)
//			//c.ws.Close()
//			c.Conn.WriteJSON(gin.H{"Message":"invalid format received"})
//			break
//		}
//		device := msg.Device
//		command := msg.Payload.Command
//		t := msg.Payload.Type
//		host := Host{}
//		//todo : command  : onhand -> Get OnHandAmount and Bind data to payload & return to Client
//		select {
//		case device == "Host" && t == "request": // นี่คือ msg type Request จาก Client
//				select {
//				case command == "cancel":
//				// model.Device.Refund  คำนวณยอดเงินที่รับและคืนเงินแยกตาม Device และแจ้ง Device ทุกตัวให้คืนเงิน
//				err := host.OrderCancel(d)
//				case command == "billing":
//				// todo: save into databse sqlite
//					res := Msg{}
//					res.Payload.Command = "billing"
//					res.Payload.Data = "Docno : xxxxxx sucessful"
//					res.Payload.Result = true
//					res.Device = "Host"
//					res.Payload.Type = "response"
//					hub.Broadcast <- res
//				case command == "onhand" && t == "request":
//					res := Msg{}
//					res.Payload.Command = "onhand"
//					res.Payload.Data = onHand.OnhandAmount
//					res.Payload.Result = true
//					res.Device = "Host"
//					res.Payload.Type = "response"
//					hub.Broadcast <- res
//				case command == "onhand" && t == "event":
//					res := Msg{}
//					res.Payload.Command = "onhand"
//
//				//ปรับยอด Onhand ตามเงินที่เข้ามา
//					amount := msg.Payload.Data
//					fmt.Println("amount : ", amount)
//
//				// todo : must be fix now - calc onHand Update
//				//onHand.OnhandAmount = onHand.OnhandAmount
//					res.Payload.Data = onHand.OnhandAmount
//					res.Payload.Result = true
//					res.Device = "Host"
//					res.Payload.Type = "response"
//					hub.Broadcast <- res
//				}
//		case device == "coin_hopper":
//				select {
//				case command = "status_changed":
//					hub.Broadcast <- msg
//				}
//		default :
//			hub.Broadcast <- msg
//		}
//	}
//}

func (c *Client) Read() {
	msg := Msg{}
	//machine := Machine{}

	for {
		err := c.Conn.ReadJSON(&msg)
		fmt.Println("command received : ", msg.Payload.Command)

		if err != nil {
			//Ghub.RemoveClient <- c
			fmt.Println("Read json from message Object Error")
			fmt.Println("Formate Not working : ", msg)
			//c.ws.Close()
			c.Conn.WriteJSON(gin.H{"message": "invalid format received"})
			break
		}

		// command check
		switch {
		// cancel order  command
		// reset ยอดเงินในตู้เป็น 0 (hardware ต้องสั่งคืนเงิน)
		case msg.Payload.Command == "cancel":
			// Reset Onhand Amount เป็น 0 และคืนเงินลูกค้า
			fmt.Println("cancel_request_starting....")
			// todo: must be send return money to Hardware
			machine.OnHand = 0
			res := Msg{}
			res.Payload.Command = "cancel"
			res.Payload.Data = "Cancel-Successful"
			res.Payload.Result = true
			res.Device = "Host"
			res.Payload.Type = "response"
			Ghub.Broadcast <- res

		// billing command ใช้สำหรับให้ Client เรียกบันทึกเข้ามาที่ Websocket
		case msg.Payload.Command == "billing":
			// todo: save into databse sqlite
			fmt.Println("billing_request_starting....")
			res := Msg{}
			res.Payload.Command = "billing"
			res.Payload.Data = "Docno : xxxxxx sucessful"
			res.Payload.Result = true
			res.Device = "Host"
			res.Payload.Type = "response"

			Ghub.Broadcast <- res

		// for Client - UI/UX call check current onhand amount
		case msg.Payload.Command == "onhand" && msg.Payload.Type == "request":
			fmt.Println("onhand_request_starting....")
			res := Msg{}
			res.Payload.Command = "onhand"
			res.Payload.Data = machine.OnHand
			res.Payload.Result = true
			res.Device = "Host"
			res.Payload.Type = "response"
			Ghub.Broadcast <- res

		// for push  totalAmount Update from hardware event and sum new onhand amount and send to UI
		case msg.Payload.Command == "onhand" && msg.Payload.Type == "event":
			res := Msg{}
			res.Payload.Command = "onhand"
			fmt.Println("onhand_event_starting....")

			// bind interface{} to i variable
			i := msg.Payload.Data

			//check type of interface{}
			//int_amount, _ := amount.(int)
			switch i.(type) {
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
			machine.OnHand = machine.OnHand + iAreaId
			fmt.Println("Current Totalamount : ", machine.OnHand)

			res.Payload.Data = machine.OnHand
			res.Payload.Result = true
			res.Device = "Host"
			res.Payload.Type = "response"
			Ghub.Broadcast <- res

		default:
			Ghub.Broadcast <- msg
		}

	}
}