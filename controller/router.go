package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/static"
	"fmt"
	"net/http"
	"github.com/mrtomyum/paybox_terminal/model"
)

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
	r.GET("/menu/:id/", GetItemsByMenuId)
	r.GET("/item/:id", GetItemById)
	//r.GET("/dev", GetDeviceIndexPage)

	r.GET("/ws", func(c *gin.Context) {
		wsPage(c.Writer, c.Request)
		fmt.Println("wsPage starting!")
	})


	// onhand initial value = 0


	return r
}

func wsPage(res http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(res, req, nil)
	fmt.Println("ws : wsPage start")
	if err != nil {
		http.NotFound(res, req)
		return
	}

	client := &model.Client{
		Conn: conn,
		Send: make(chan model.Msg),
	}

	Ghub.AddClient <- client

	go client.Write()
	go client.Read()
}
