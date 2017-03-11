package ctrl

import (
	"fmt"
	"net/http"
	"github.com/mrtomyum/paybox_web/model"
	"log"
	"net/url"
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
	c := &model.Client{
		Ws:   conn,
		Send: make(chan *model.Message),
		Name: "web",
		Msg:  &model.Message{},
	}
	//fmt.Println("Web:", c.Name, "...start send <-c to model.H.Webclient")
	model.H.Web = c
	fmt.Println("Start WebSocket connection from Web:", conn.RemoteAddr())
	go c.Write()
	c.Read() // ดัก Event message ที่จะส่งมาตอนไหนก็ไม่รู้
}

// CallDev() เพื่อให้โปรแกรม Host เรียก WebSocket ไปยัง HW_SERVICE ที่พอร์ท 9999
// ใช้สั่งงาน Request และรับ Event/Response จาก Device ต่างๆ
func CallDev() {
	//u := url.URL{Scheme: "ws", Host: "127.0.0.1:9999", Path: "/"}
	u := url.URL{Scheme:"ws", Host:"192.168.10.64:9999", Path: "/"}
	log.Printf("connecting to %s", u.String())
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Println("Error call HW_SERVICE Websocket:", err)
		return
	}
	defer conn.Close()
	c := &model.Client{
		Ws:   conn,
		Send: make(chan *model.Message),
		Name: "dev",
		Msg:  &model.Message{},
	}
	model.H.Hw = c
	fmt.Println("Start Websocket to HW_SERVICE connected:", conn.RemoteAddr())
	go c.Write()
	c.Read()
}

func ServHW(w http.ResponseWriter, r *http.Request) {
	fmt.Println("start ServDev Websocket for Device...")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil { // pass this func if currently no WebSocket service
		fmt.Println(err)
		return
	}
	defer conn.Close()
	//fmt.Println("start New Device connection success...")
	c := &model.Client{
		Ws:   conn,
		Send: make(chan *model.Message),
		Name: "dev",
	}
	fmt.Println("Start Dev Connection:")
	model.H.Hw = c
	go c.Write()
	c.Read()
}

