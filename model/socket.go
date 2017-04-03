package model

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/labstack/gommon/log"
)

type Socket struct {
	Name string
	Conn *websocket.Conn
	Msg  *Message
	Send chan *Message
}

type Message struct {
	Device  string `json:"device"`
	Type    string `json:"type"`
	Command string `json:"command"`
	Result  bool `json:"result,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (s *Socket) Read() {
	defer func() {
		s.Conn.Close()
	}()

	m := &Message{}
	for {
		err := s.Conn.ReadJSON(&m)
		if err != nil {
			log.Debug(s.Name, "<<===Conn.ReadJSON Error on:", err)
			break
		}
		log.Debug("<===*Socket.ReadJSON====", s.Name, s.Conn.RemoteAddr(), m)
		s.Msg = m

		switch {
		case s.Name == "UI":
			//fmt.Println("Read::Web UI Connection message")
			s.onUiEvent()
		case s.Name == "HW":
			//fmt.Println("Read::Device Connection message")
			s.onHwEvent()
		}
	}
}

// Write() ส่วนเขียนข้อความไปยัง WebSocket
func (s *Socket) Write() {
	log.Info("###*Socket.Write()### START ###", s.Name, s.Conn.RemoteAddr())
	defer log.Info("###*Socket.Write()### END ###", s.Name, s.Conn.RemoteAddr())
	defer s.Conn.Close()
	for {
		select {
		case m, ok := <-s.Send:
			if !ok {
				s.Conn.WriteJSON(gin.H{"message": "Cannot send data"})
				log.Debug("===>>>lose WS connection:", s.Conn.RemoteAddr())
				return
			}
			s.Conn.WriteJSON(m)
			log.Debugf("\n===*Socket.WriteJSON===> %s:%v %v\n", s.Name, s.Conn.RemoteAddr(), m)
		}
	}
}

// onUiEvent แยกเส้นทาง Message Request จาก Web Frontend โดยแยกตาม Command ต่างๆ
func (s *Socket) onUiEvent() {
	// ปกติแล้ว  Web จะไม่สั่งการ Device ตรงๆ แต่จะสั่งผ่าน Host ให้ Host ทำงานระดับล่างแทน
	// แต่ตรงนี้มีไว้สำหรับการ Debug ผ่าน Web GUI fmt.Println("Request message from Web")
	switch s.Msg.Command {
	case "onhand":
		PM.sendOnHand(s)
	case "cancel":
		PM.Cancel(s)
	default:
		log.Debug("onUiEvent(): default: Unknown Command for web client=>", s.Msg.Command)
	}
}

// onHwEvent เป็นการแยกเส้นทาง Message จาก Device Event และ Response
// โดย Function นี้จะแยก message ตาม Device ก่อน แล้วจึงแยกเส้นทางตาม Command
func (s *Socket) onHwEvent() {
	//fmt.Println("onHwEvent():", s.Msg)
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
		log.Debug("event cannot find function/message=", s.Msg)
	}
}
