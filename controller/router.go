package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/gorilla/websocket"
	"fmt"
)

func Router(r *gin.Engine) *gin.Engine {
	r.GET("/", GetMenuIndex)
	r.GET("/item/:id", GetItemByMenu)
	r.POST("/order", PostNewOrderSub)
	r.DELETE("/order", DeleteOrder)
	r.DELETE("/order/item/:line", DeleteOrderItem)

	r.GET("/ws", func(c *gin.Context) {
		wsView(c.Writer, c.Request)
	})
	return r
}

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wsView(w http.ResponseWriter, r *http.Request) {
	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}
	go func() {
		for {
			t, msg, err := conn.ReadMessage()
			if err != nil {
				break
			}
			//select {
			//
			//}
			conn.WriteMessage(t, msg)
		}
	}()
}