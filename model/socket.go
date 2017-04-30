package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
)

type Socket struct {
	Name string
	Conn *websocket.Conn
	Msg  *Message
	Send chan *Message
}

type Message struct {
	Device  string      `json:"device"`
	Type    string      `json:"type"`
	Command string      `json:"command"`
	Result  bool        `json:"result,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (s *Socket) Read(done chan bool) {
	defer func() {
		s.Conn.Close()
		done <- true
	}()

	m := &Message{}
	//count := 0
	fmt.Println("###*Socket.Read()### START ###", s.Name, s.Conn.RemoteAddr())
	for {
		err := s.Conn.ReadJSON(&m)
		if err != nil {
			log.Println(s.Name, "<<===Conn.ReadJSON Error on:", err)
			break
		}
		log.Println("<===*Socket.ReadJSON====", s.Name, s.Conn.RemoteAddr(), m)
		s.Msg = m

		// Detect Ghost Message!!
		// ดัก device = "bill_acc" || "coin_acc"
		// ถ้า payment.isOpen() false  ให้ bypass continue

		switch s.Name {
		case "UI":
			//log.Println("*Socket.Read::Web UI Connection message 2")
			go s.onUiEvent()
		case "HW":
			//log.Println("*Socket.Read::Detectevice Connection message 2")
			go s.onHwEvent()
		default:
			log.Println("*Socket.Read::Unknown message", s.Msg)
		}
	}
}

// Write() ส่วนเขียนข้อความไปยัง WebSocket
func (s *Socket) Write() {
	log.Println("###*Socket.Write()### START ###", s.Name, s.Conn.RemoteAddr())
	defer log.Println("###*Socket.Write()### END ###", s.Name, s.Conn.RemoteAddr())
	defer s.Conn.Close()
	for {
		select {
		case m, ok := <-s.Send:
			if !ok {
				s.Conn.WriteJSON(gin.H{"message": "Cannot billCh data"})
				log.Println("===>>>lose WS connection:", s.Conn.RemoteAddr())
				return
			}
			s.Conn.WriteJSON(m)
			log.Printf("\n===*Socket.WriteJSON===> %s:%v %v\n", s.Name, s.Conn.RemoteAddr(), m)
		}
	}
}

// onUiEvent แยกเส้นทาง Message Request จาก Web Frontend โดยแยกตาม Command ต่างๆ
// ปกติแล้ว  Web จะไม่สั่งการ Device ตรงๆ แต่จะสั่งผ่าน Host ให้ Host ทำงานระดับล่างแทน
// แต่ตรงนี้มีไว้สำหรับการ Debug ผ่าน Web GUI fmt.Println("Request message from Web")
func (s *Socket) onUiEvent() {
	switch s.Msg.Command {
	case "onhand":
		PM.sendOnHand(s)
	case "cancel":
		PM.Cancel()
		//case "order":
		//	go PM.Order(s)
	default:
		log.Println("onUiEvent(): default: Unknown Command for web client=>", s.Msg.Command)
	}
}

// onHwEvent เป็นการแยกเส้นทาง Message จาก Device Event และ Response
// โดย Function นี้จะแยก message ตาม Device ก่อน แล้วจึงแยกเส้นทางตาม Command
func (s *Socket) onHwEvent() {
	switch s.Msg.Device {
	case "coin_hopper":
		CH.event(s)
	case "coin_acc":
		CA.event(s)
	case "bill_acc":
		BA.event(s)
	case "printer":
		P.event(s)
	case "mainboard":
		MB.event(s)
	default:
		log.Println("event cannot find function/message=", s.Msg)
	}
}
