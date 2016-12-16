package controller

//import
import (
	"github.com/gorilla/websocket"
	"github.com/mrtomyum/paybox_terminal/model"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool {
		return true
	},
}

var machine = model.Machine{
	Id:     "1",
	OnHand: 0,
}



//type Client struct {
//	ws   *websocket.Conn
//	send chan model.Msg
//}
//
//func (c *Client) write() {
//	defer func() {
//		c.ws.Close()
//	}()
//
//	for {
//		select {
//		case msg, ok := <-c.send:
//			if !ok {
//				//  c.ws.WriteMessage(websocket.CloseMessage, []byte{})
//				c.ws.WriteJSON(gin.H{"message": "Connot to Send data"})
//				return
//			}
//
//		//c.ws.WriteMessage(websocket.TextMessage, message)
//			c.ws.WriteJSON(msg)
//		}
//	}
//}
//
//func (c *Client) read() {
//	msg := model.Msg{}
//	for {
//		err := c.ws.ReadJSON(&msg)
//		fmt.Println("command received : ", msg.Payload.Command)
//
//		if err != nil {
//			//Ghub.RemoveClient <- c
//			fmt.Println("Read json from message Object Error")
//			fmt.Println("Formate Not working : ", msg)
//			//c.ws.Close()
//			c.ws.WriteJSON(gin.H{"message": "invalid format received"})
//			break
//		}
//
//		// command check
//		switch {
//		// cancel order  command
//		// reset ยอดเงินในตู้เป็น 0 (hardware ต้องสั่งคืนเงิน)
//		case msg.Payload.Command == "cancel":
//			// Reset Onhand Amount เป็น 0 และคืนเงินลูกค้า
//			fmt.Println("cancel_request_starting....")
//			// todo: must be send return money to Hardware
//			machine.OnHand = 0
//			res := model.Msg{}
//			res.Payload.Command = "cancel"
//			res.Payload.Data = "Cancel-Successful"
//			res.Payload.Result = true
//			res.Device = "Host"
//			res.Payload.Type = "response"
//			Ghub.Broadcast <- res
//
//		// billing command ใช้สำหรับให้ Client เรียกบันทึกเข้ามาที่ Websocket
//		case msg.Payload.Command == "billing":
//			// todo: save into databse sqlite
//			fmt.Println("billing_request_starting....")
//			res := model.Msg{}
//			res.Payload.Command = "billing"
//			res.Payload.Data = "Docno : xxxxxx sucessful"
//			res.Payload.Result = true
//			res.Device = "Host"
//			res.Payload.Type = "response"
//
//			Ghub.Broadcast <- res
//
//		// for Client - UI/UX call check current onhand amount
//		case msg.Payload.Command == "onhand" && msg.Payload.Type == "request":
//			fmt.Println("onhand_request_starting....")
//			res := model.Msg{}
//			res.Payload.Command = "onhand"
//			res.Payload.Data = machine.OnHand
//			res.Payload.Result = true
//			res.Device = "Host"
//			res.Payload.Type = "response"
//			Ghub.Broadcast <- res
//
//		// for push  totalAmount Update from hardware event and sum new onhand amount and send to UI
//		case msg.Payload.Command == "onhand" && msg.Payload.Type == "event":
//			res := model.Msg{}
//			res.Payload.Command = "onhand"
//			fmt.Println("onhand_event_starting....")
//
//			// bind interface{} to i variable
//			i := msg.Payload.Data
//
//			//check type of interface{}
//			//int_amount, _ := amount.(int)
//			switch i.(type) {
//			case float64:
//				fmt.Println("amnount type : float64")
//			case float32:
//				fmt.Println("amnount type : float32")
//			case int64:
//				fmt.Println("amnount type : int64")
//			}
//
//			// convert interface{} to int
//			// example onhand_event
//			var iAreaId int = int(i.(float64))
//
//			// Update Current OnHand TOTAL
//			machine.OnHand = machine.OnHand + iAreaId
//			fmt.Println("Current Totalamount : ", machine.OnHand)
//
//			res.Payload.Data = machine.OnHand
//			res.Payload.Result = true
//			res.Device = "Host"
//			res.Payload.Type = "response"
//			Ghub.Broadcast <- res
//
//		default:
//			Ghub.Broadcast <- msg
//		}
//
//	}
//}
