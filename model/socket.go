package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"encoding/json"
	"os"
	"os/signal"
)

type Socket struct {
	Conn *websocket.Conn
	Send chan *Message
	Name string
	Msg  *Message
}

type Message struct {
	Device  string `json:"device"`
	Type    string `json:"type"`
	Command string `json:"command"`
	Result  bool `json:"result,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (c *Socket) Read() {
	done := make(chan struct{})
	defer c.Conn.Close()
	defer close(done)

	m := &Message{}
	for {
		err := c.Conn.ReadJSON(&m)
		if err != nil {
			log.Println("Conn.ReadJSON Error on : ", c.Name, " :", err)
			break
		}
		c.Msg = m
		// Debug check json Encode
		b, err := json.Marshal(m)
		if err != nil {
			fmt.Println("error:", err)
		}
		os.Stdout.Write(b)
		fmt.Println("Socket", c.Name, " read JSON message. Command:", m.Command)

		switch {
		case c.Name == "web":
			//fmt.Println("Read::Web UI Connection message")
			c.WebEvent()
		case c.Name == "dev":
			//fmt.Println("Read::Device Connection message")
			c.HwEvent()
		default:
			fmt.Println("Error: Case default: Message==>", m)
			m.Type = "response"
			m.Data = "Hello"
			c.Send <- m
		}
	}
}

func (c *Socket) Write() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	fmt.Println("=======*Socket.Write()========")
	defer fmt.Println("=====*Socket.Write()=== END ====")
	defer c.Conn.Close()
	for {
		select {
		case m, ok := <-c.Send:
			if !ok {
				c.Conn.WriteJSON(gin.H{"message": "Cannot send data"})
				return
			}
			c.Conn.WriteJSON(m)
			// Debug check json Encode
			b, err := json.Marshal(m)
			if err != nil {
				fmt.Println("error:", err)
			}
			os.Stdout.Write(b)
		case <-interrupt:
			log.Println("interrupt")
			// To cleanly close a connection, a client should send a close
			// frame and wait for the server to close the connection.
			err := c.Conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			return
		}
	}
}

// WebEvent แยกเส้นทาง Message Request จาก Web Frontend โดยแยกตาม Command ต่างๆ
func (c *Socket) WebEvent() {
	// ปกติแล้ว  Web จะไม่สั่งการ Device ตรงๆ
	// จะสั่งผ่าน Host ให้ Host ทำงานระดับล่างแทน
	// แต่ตรงนี้มีไว้สำหรับการ Debug ผ่าน Web GUI
	fmt.Println("Request message from Web")
	switch c.Msg.Command {
	case "onhand":
		PM.sendOnHand(c)
	case "cancel":
		PM.Cancel(c)
	default:
		log.Println("WebEvent(): default: Unknown Command for web client=>", c.Msg.Command)
	}
}

// HwEvent เป็นการแยกเส้นทาง Message จาก Device Event และ Response
// โดย Function นี้จะแยก message ตาม Device ก่อน แล้วจึงแยกเส้นทางตาม Command
func (c *Socket) HwEvent() {
	fmt.Println("HwEvent():", c.Msg)
	switch c.Msg.Device {
	case "coin_hopper":
		CH.Event(c)
	case "coin_acc":
		CA.Event(c)
	case "bill_acc":
		BA.Event(c)
	case "printer":
		P.Event(c)
	case "mainboard":
		M.Event(c)
	default:
		log.Println("event cannot find function/message=", c.Msg)
	}
}

