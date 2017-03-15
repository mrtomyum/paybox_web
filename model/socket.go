package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	//"encoding/json"
	//"os"
	//"os/signal"
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
	defer func() {
		c.Conn.Close()
	}()

	m := &Message{}
	for {
		err := c.Conn.ReadJSON(&m)
		fmt.Println("<========*Socket.Read()==========", c.Name, c.Conn.RemoteAddr(), m)
		if err != nil {
			log.Println(c.Name, "<<===Conn.ReadJSON Error on:", err)
			break
		}
		c.Msg = m

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

// Write() ส่วนเขียนข้อความไปยัง WebSocket
func (c *Socket) Write() {
	//interrupt := make(chan os.Signal, 1)
	//signal.Notify(interrupt, os.Interrupt)

	fmt.Println("=======*Socket.Write()== START =>", c.Name, c.Conn.RemoteAddr())
	defer fmt.Println("=====*Socket.Write()=== END ==>", c.Name, c.Conn.RemoteAddr())
	defer c.Conn.Close()
	for {
		select {
		case m, ok := <-c.Send:
			if !ok {
				c.Conn.WriteJSON(gin.H{"message": "Cannot send data"})
				log.Println("===>>>lose WS connection:", c.Conn.RemoteAddr())
				return
			}
			c.Conn.WriteJSON(m)
			fmt.Printf("\n====*Socket.Conn.WriteJSON()====> %s:%v\nMessage:========\n%vData:========\n%v\n", c.Name, c.Conn.RemoteAddr(), m, m.Data)
		}
	}
}

// WebEvent แยกเส้นทาง Message Request จาก Web Frontend โดยแยกตาม Command ต่างๆ
func (c *Socket) WebEvent() {
	// ปกติแล้ว  Web จะไม่สั่งการ Device ตรงๆ แต่จะสั่งผ่าน Host ให้ Host ทำงานระดับล่างแทน
	// แต่ตรงนี้มีไว้สำหรับการ Debug ผ่าน Web GUI fmt.Println("Request message from Web")
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
	//fmt.Println("HwEvent():", c.Msg)
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
		MB.Event(c)
	default:
		log.Println("event cannot find function/message=", c.Msg)
	}
}
