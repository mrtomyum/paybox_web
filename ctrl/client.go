package ctrl

import (
	"fmt"
	"net/http"
	"github.com/mrtomyum/paybox_terminal/model"
	"log"
	"net/url"
	"github.com/gorilla/websocket"
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

func CallDev() error {
	u := url.URL{Scheme:"ws", Host:"localhost:9999", Path: "/ws"}
	log.Printf("connecting to %s", u.String())
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return err
	}
	defer conn.Close()
	c := &model.Client{
		Ws:   conn,
		Send: make(chan *model.Message),
		Name: "web",
		Msg:  &model.Message{},
	}
	go c.Write()
	c.Read()
	return nil
}
