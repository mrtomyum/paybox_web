package controller
//
//import (
//	"net/http"
//	"fmt"
//	"github.com/mrtomyum/paybox_terminal/model"
//)

////  Mock Web socket server for host at port 9999===>Not active in production.
//func WsDevice(w http.ResponseWriter, r *http.Request) {
//	conn, err := upgrader.Upgrade(w, r, nil)
//	fmt.Println("ws : wsDevice start")
//	if err != nil {
//		http.NotFound(w, r)
//		return
//	}
//	defer conn.Close()
//	device := model.Device{
//		Conn: conn,
//		Send: make(chan model.Msg),
//	}
//
//	// Listening to Event from server
//	go device.Write()
//	device.Read()
//
//}

// Web socket server waiting for Web Front end connection.
//func wsServer(w http.ResponseWriter, r *http.Request) {
//	conn, err := wsUpgrader.Upgrade(w, r, nil)
//	if err != nil {
//		fmt.Println("Failed to set websocket upgrade: %+v", err)
//		http.NotFound(w, r)
//		return
//	}
//
//	msg := model.Msg{}
//	onHand := model.Machine{}
//	//	event := make(chan string, 10)
//
//	go func() {
//		for {
//			err := conn.ReadJSON(&msg)
//			if err != nil {
//				log.Println("Error Read JSON:", msg)
//			}
//			data := model.Money{}
//			//json.Unmarshal(msg, &data)
//
//			switch msg.Payload.Command {
//			case "onhand":
//			case "cancel":
//			case "billing":
//
//
//
//
//			}
//			if data.Job == "onHand" {
//				// return onhand
//				//				onHand.Job = "onHand"
//				conn.WriteJSON(msg)
//
//			}
//			if data.Job == "money" {
//				// เติมเงินเข้าตู้
//
//				fmt.Println("Amount:", data.Amount)
//				fmt.Println("job :", data.Job)
//
//				// Update Onhand amount
//				onHand.OnHand = onHand.OnHand + data.Amount
//				conn.WriteJSON(msg)
//				fmt.Println("ON Hand Amount : ", onHand.OnHand)
//			}
//
//			if data.Job == "print" {
//				// เติมเงินเข้าตู้
//				onHand.OnHand = 0
//
//				// connect to ws printer board
//				//msg, err = json.Marshal(gin.H{"message":"success", "Job":"print"})
//				msg.Device = "printer"
//				msg.Payload.Type = "response"
//				msg.Payload.Result = true
//				msg.Payload.Command = "print"
//				conn.WriteJSON(msg)
//				fmt.Println("ON Hand Amount : ", onHand.OnHand)
//			}
//
//			log.Println(msg)
			//			e := <-event
			//			 if e != nil {
			//				msg.Payload.Command = e
			//				conn.WriteJSON(msg)
			//			}

//		}
//	}()
//}

