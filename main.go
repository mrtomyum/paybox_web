package main

import (
	"github.com/gin-gonic/gin"
	c "github.com/mrtomyum/paybox_terminal/controller"
	"net/http"
	"github.com/gorilla/websocket"
	"github.com/labstack/gommon/log"
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


func wshandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}
	go func() {
		for {
			t, msg, err := conn.ReadMessage()
			//log.Print(string(msg))

			fmt.Println("received data :", string(msg))

			if err != nil {
				break
			}

			msg, err = json.Marshal(gin.H{"message":string(msg)})
			if err != nil {
				log.Print("Eror Marshal gin.H")
			}
			//conn.WriteMessage(t, msg)
			conn.WriteMessage(t, msg)

		}
	}()
}