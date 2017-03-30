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

// ServWeb ทำงานเมื่อ Web Socket เรียกเพจ /ws ระบบ Host จะทำตัวเป็น Server ให้ Socket จาก WebUI เชื่อมต่อเข้ามาคุย
func ServWeb(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("start ServWeb Websocket for Web...")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	s := &model.Socket{
		Name: "UI",
		Conn: conn,
		Send: make(chan *model.Message),
		Msg:  &model.Message{},
	}
	model.H.Web = s
	fmt.Println("Open WebSocket from Web:", conn.RemoteAddr())
	go s.Write()
	s.Read() // ดัก Event message ที่จะส่งมาตอนไหนก็ไม่รู้
}

// OpenSocket() เพื่อให้โปรแกรม Host เรียก WebSocket ไปยัง HW_SERVICE ที่พอร์ท 9999
// ใช้สั่งงาน Request และรับ Event/Response จาก Device ต่างๆ
func OpenSocket() {
	//u := url.URL{Scheme: "ws", Host: "127.0.0.1:9999", Path: "/"}
	u := url.URL{Scheme:"ws", Host:"192.168.10.64:9999", Path: "/"}
	log.Printf("connecting to %s", u.String())
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("Error call HW_SERVICE Websocket:", err)
	}
	s := &model.Socket{
		Send: make(chan *model.Message),
		Name: "HW",
		Msg:  &model.Message{},
	}
	fmt.Printf("Open Websocket to %v connected: %v\n", s.Name, conn.RemoteAddr())
	model.H.Hw = s
	s.Conn = conn
	defer conn.Close()
	go s.Write()
	s.Read()
}


