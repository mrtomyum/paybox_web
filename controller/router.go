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
	//r.Static("/public", "./view/public")
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

// wsPage ทำงานเมื่อ Web Client เรียกเพจ /ws ระบบ Host จะทำตัวเป็น
// Server ให้ Client เชื่อมต่อเข้ามา รัน goroutine จาก client.Write() & .Read()
func wsPage(res http.ResponseWriter, req *http.Request) {
	conn1, err := upgrader.Upgrade(res, req, nil)
	fmt.Println("ws : wsPage start")
	if err != nil {
		http.NotFound(res, req)
		return
	}
	defer conn1.Close()
	client := &model.Client{
		Conn: conn1,
		Send: make(chan model.Msg),
	}
	model.MyHub.AddClient <- client
	go client.Write()
	client.Read()

	// Dial conn2 to Device Websocket
	//url := "http://localhost:9999/ws"
	//conn2, _, err := websocket.DefaultDialer.Dial(url, nil)
	//if err != nil {
	//	log.Println("dial:", err)
	//}
	//defer conn2.Close()
	//device := model.Device{}
	//go device.Write()
	//go device.Read()
}

