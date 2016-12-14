package controller

import (
	"net/http"
	"github.com/gorilla/websocket"
	"github.com/mrtomyum/paybox_terminal/model"
//"net/url"
	"log"
	"fmt"
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Web socket client for Web Front end.
//func WsClient() {
//	addr := "localhost:9999"
//	u := url.URL{Scheme:"ws", Host: addr, Path: "/ws"}
//	log.Printf("กำลังเชื่อมต่อไปที่ %s", u.String())
//
//	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
//	if err != nil {
//		log.Fatal("dial:", err)
//	}
//	defer conn.Close()
//
//	m := model.Msg{}
//	// Listening to Event from server
//	go func() {
//		defer conn.Close()
//		for {
//			err := conn.ReadJSON(&m)
//			if err != nil {
//				log.Println("read:", err)
//				break
//			}
//
//			switch m.Device {
//			case "coin_hopper":
//				// implementing
//				ch := model.CoinHopper{}
//				ch.CheckMsg()
//			case "coin_acc":
//			case "bill_acc":
//			case "printer":
//			}
//			conn.WriteJSON(&m)
//		}
//	}()
//
//}

// Web socket server waiting for Web Front end connection.
func wsServer(w http.ResponseWriter, r *http.Request) {
	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		http.NotFound(w, r)
		return
	}

	msg := model.Msg{}
	onHand := model.Machine{}
	//	event := make(chan string, 10)

	go func() {
		for {
			err := conn.ReadJSON(&msg)
			if err != nil {
				log.Println("Error Read JSON:", msg)
			}
			data := model.Money{}
			//json.Unmarshal(msg, &data)

			switch msg.Payload.Command {
			case "onhand":
			case "cancel":
			case "billing":




			}
			if data.Job == "onHand" {
				// return onhand
				//				onHand.Job = "onHand"
				conn.WriteJSON(msg)

			}
			if data.Job == "money" {
				// เติมเงินเข้าตู้

				fmt.Println("Amount:", data.Amount)
				fmt.Println("job :", data.Job)

				// Update Onhand amount
				onHand.OnHand = onHand.OnHand + data.Amount
				conn.WriteJSON(msg)
				fmt.Println("ON Hand Amount : ", onHand.OnHand)
			}

			if data.Job == "print" {
				// เติมเงินเข้าตู้
				onHand.OnHand = 0

				// connect to ws printer board
				//msg, err = json.Marshal(gin.H{"message":"success", "Job":"print"})
				msg.Device = "printer"
				msg.Payload.Type = "response"
				msg.Payload.Result = true
				msg.Payload.Command = "print"
				conn.WriteJSON(msg)
				fmt.Println("ON Hand Amount : ", onHand.OnHand)
			}

			log.Println(msg)
			//			e := <-event
			//			 if e != nil {
			//				msg.Payload.Command = e
			//				conn.WriteJSON(msg)
			//			}

		}
	}()
}

