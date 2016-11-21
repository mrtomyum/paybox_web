package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/static"
	"net/http"
	"github.com/gorilla/websocket"
	"fmt"
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Router(r *gin.Engine) *gin.Engine {
	r.LoadHTMLGlob("view/**/*.tpl")
	r.Static("/html", "./view/html")
	r.Static("/public", "./view/public")
	r.Static("/js", "./view/public/js")
	r.Static("/css", "./view/public/css")
	r.Static("/img", "./view/public/img")
	r.Static("/json", "./view/public/json")

	//r.GET("/", GetIndex)
	r.Use(static.Serve("/", static.LocalFile("view", true)))
	r.GET("/menu", GetMenu)
	r.GET("/menu/:id/", GetItemByMenuId)
	r.GET("/item/:id", GetItemByMenu)
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
