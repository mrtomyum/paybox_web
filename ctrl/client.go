package ctrl

import (
	"fmt"
	"net/http"
	"github.com/mrtomyum/paybox_terminal/model"
	"log"
	"time"
	"os"
	"os/signal"
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

	// Dial to Device WS server
	err = CallDev()
	if err != nil {
		log.Println("Error call Device Websocket:", err)
	}

	fmt.Println("Start Web connection")
	go c.Write()
	c.Read() // ดัก Event message ที่จะส่งมาตอนไหนก็ไม่รู้
}

func CallDev() error {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	u := url.URL{Scheme:"ws", Host:"localhost:9999", Path: "/ws"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()
	done := make(chan struct{})
	m := &model.Message{}

	go func() {
		defer c.Close()
		defer close(done)
		for {
			err := c.ReadJSON(&m)
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Panicf("recev: %s", m)
		}
	}()
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case t := <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("write:", err)
				return err
			}
		case <-interrupt:
			log.Println("interrupt")
			// To cleanly close a connection, a client should send a close
			// frame and wait for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return err
			}
				select {
				case <-done:
				case <-time.After(time.Second):
				}
			c.Close()
			return nil
		}
	}
	return nil
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

