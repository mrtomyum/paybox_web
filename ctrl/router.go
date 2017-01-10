package ctrl

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/static"
	"fmt"
	"net/http"
	"github.com/mrtomyum/paybox_terminal/model"
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
	//r.LoadHTMLGlob("view/**/*.tpl")
	//r.Static("/html", "./view/html")
	r.Static("/js", "./view/public/js")
	r.Static("/css", "./view/public/css")
	r.Static("/img", "./view/public/img")
	r.Static("/json", "./view/public/json")

	//r.GET("/", GetIndex)
	r.Use(static.Serve("/", static.LocalFile("view", true)))
	r.GET("/menu", GetMenu)
	r.GET("/menu/:id/", GetItemsByMenuId)
	r.GET("/item/:id", GetItemById)

	r.GET("/web", func(c *gin.Context) {
		ServWeb(c.Writer, c.Request)
	})
	r.GET("/dev", func(c *gin.Context) {
		ServDev(c.Writer, c.Request)
	})
	return r
}

// wsServer ทำงานเมื่อ Web Client เรียกเพจ /ws ระบบ Host จะทำตัวเป็น
// Server ให้ Client เชื่อมต่อเข้ามา รัน goroutine จาก client.Write() & .Read()

func ServWeb(w http.ResponseWriter, r *http.Request) {
	fmt.Println("start ServWeb Websocket for Web...")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	fmt.Println("start New Web connection success...")
	c := &model.Client{
		Ws:   conn,
		Send: make(chan *model.Message),
		Name: "web",
	}
	fmt.Println("Web:", c.Name, "...start send <-c to model.H.Webclient")
	model.H.Web = c
	fmt.Println("start go c.Write()")
	go c.Write()
	c.Read() // ดัก Event message ที่จะส่งมาตอนไหนก็ไม่รู้
}

func ServDev(w http.ResponseWriter, r *http.Request) {
	fmt.Println("start ServDev Websocket for Device...")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	fmt.Println("start New Device connection success...")
	c := &model.Client{
		Ws:   conn,
		Send: make(chan *model.Message),
		Name: "dev",
	}
	fmt.Println("Dev:", c.Name)
	model.H.Dev = c
	go c.Write()
	c.Read()
}