package main

import (
	"github.com/gin-gonic/gin"
	c "github.com/mrtomyum/paybox_terminal/controller"
	"net/http"
	"github.com/gorilla/websocket"
//"github.com/labstack/gommon/log"
	"fmt"
	"encoding/json"
)

func main() {
	r := gin.Default()
	app := c.Router(r)


	r.GET("/ws", func(c *gin.Context) {
		wshandler(c.Writer, c.Request)
	})

	app.Run(":8888")
}
var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}



type Money struct {
	Job    string
	Amount int
}

type Onhand struct {
	OnhandAmount int
}

var Remain = Onhand{}

//Remain.OnhandAmount = 0

func wshandler(w http.ResponseWriter, r *http.Request) {

	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}
	go func() {
		for {
			t, msg, err := conn.ReadMessage()
			data := Money{}
			json.Unmarshal(msg, &data)


			if data.Job == "onHand" {
				// return onhand
				msg, err = json.Marshal(Remain)
				conn.WriteMessage(t, msg)

			}
			if data.Job == "money" {
				// เพิ่ม Onhand


				fmt.Println("Amount:", data.Amount)
				fmt.Println("job :", data.Job)

				Remain.OnhandAmount = Remain.OnhandAmount + data.Amount
				conn.WriteMessage(t, msg)
				fmt.Println("ON Hand Amount : ", Remain.OnhandAmount)
			}




			//log.Print(string(msg))

			fmt.Println("received data :", string(msg))

			if err != nil {
				break
			}

			//			msg, err = json.Marshal(gin.H{"message":string(msg)})
			//			if err != nil {
			//				log.Print("Eror Marshal gin.H")
			//			}
			//conn.WriteMessage(t, msg)
			//			conn.WriteMessage(t, msg)

		}
	}()
}