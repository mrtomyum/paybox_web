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
		wsServer(c.Writer, c.Request)
		fmt.Println("wsPage starting!")
		//WsDevice(c.Writer, c.Request)
		//fmt.Println("wsDevice starting")
	})
	return r
}

// wsServer ทำงานเมื่อ Web Client เรียกเพจ /ws ระบบ Host จะทำตัวเป็น
// Server ให้ Client เชื่อมต่อเข้ามา รัน goroutine จาก client.Write() & .Read()
func wsServer(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	fmt.Println("ws : wsServer start")
	if err != nil {
		http.NotFound(w, r)
		return
	}
	defer conn.Close()

	// client จะต้องระบุตัวตนว่าเป็น "dev" หรือ "web" เข้ามาใน header "name"
	clientName := r.Header.Get("name")
	c := &model.Client{
		Conn: conn,
		Msg: make(chan model.Msg),
		Name: clientName,
	}
	11
	dClient <- c
	go c.Write()
	c.Read()

	// Dial conn2 to Device Websocket
	//url := "http://localhost:9999/ws"
	//devConn, _, err := websocket.DefaultDialer.Dial(url, nil)
	//if err != nil {
	//	log.Println("dial:", err)
	//}
	//defer devConn.Close()
	//device := model.Device{
	//	Conn: devConn,
	//	Send: make(chan model.Msg),
	//}
	//go device.Write()
	//device.Read()
}

