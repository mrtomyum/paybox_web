package ctrl

import (
	"fmt"
	"net/http"
	"github.com/mrtomyum/paybox_terminal/model"
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
	fmt.Println("start New Web connection success...")

	c := &model.Client{
		Ws:   conn,
		Send: make(chan *model.Message),
		Name: "web",
		Msg:  &model.Message{},
	}
	fmt.Println("Web:", c.Name, "...start send <-c to model.H.Webclient")
	model.H.Web = c
	fmt.Println("Start Web connection")
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
	fmt.Println("Start Dev Connection:")
	model.H.Dev = c
	go c.Write()
	c.Read()
}

func CallDev() {
	//u := url.URL{Scheme:"ws", Host:"127.0.0.1:9999", Path: "/"}
	u := url.URL{Scheme:"ws", Host:"192.168.10.64:9999", Path: "/"}
	log.Printf("connecting to %s", u.String())
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Println("Error call Device Websocket:", err)
	}
	fmt.Println("connected", conn)
	defer conn.Close()
	c := &model.Client{
		Ws:   conn,
		Send: make(chan *model.Message),
		Name: "dev",
		Msg:  &model.Message{},
	}
	model.H.Dev = c
	fmt.Println("Start Dev Connection, model.H.Dev:", model.H.Dev)
	go c.Write()
	c.Read()
}
