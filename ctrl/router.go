package ctrl

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/static"
	"net/http"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)
func Router(r *gin.Engine) *gin.Engine {
	// for Static HTML
	r.LoadHTMLGlob("view/**/*.tpl")
	//r.Static("/html", "./view/html")
	r.Static("/js", "./view/public/js")
	r.Static("/css", "./view/public/css")
	r.Static("/img", "./view/public/img")
	r.Static("/json", "./view/public/json")

	// for Web endpoint call data in JSON
	r.Use(static.Serve("/", static.LocalFile("view", true)))
	r.GET("/menu", GetMenu)
	r.GET("/menu/:id/", GetItemsByMenuId)
	r.GET("/item/:id", GetItemById)
	r.POST("/sale", NewSale)

	// for WebSocket Connection
	r.GET("/web", func(c *gin.Context) {
		ServWeb(c.Writer, c.Request)
	})
	r.GET("/dev", func(c *gin.Context) {
		ServDev(c.Writer, c.Request)
	})
	return r
}

