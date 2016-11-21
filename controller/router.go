package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/gorilla/websocket"
	"fmt"
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Router(r *gin.Engine) *gin.Engine {
	r.GET("/", GetIndex)
	r.GET("/item/:id", GetItemByMenu)
	r.GET("/list", GetMenu)
	r.GET("/menu/:id/", GetItemByMenuId)
	r.GET("/dev", GetDeviceIndexPage)

	r.GET("/view", func(c *gin.Context) {
		wsView(c.Writer, c.Request)
	})
	return r
}

// Test websocket server --not in production--
func wsView(w http.ResponseWriter, r *http.Request) {
	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Failed to set websocket upgrade: %+v", err)
		return
	}
	m := Msg{}
	go func() {
		for {
			err := conn.ReadJSON(&m)
			if err != nil {
				break
			}
			if m.Topic == "ping" {
				m = Msg{Topic:"pong", Status: OK}
			}
			conn.WriteJSON(&m)
		}
	}()
}
