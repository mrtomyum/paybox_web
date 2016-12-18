package model

import (
	"log"
	"github.com/gorilla/websocket"
	"github.com/gin-gonic/gin"
	"fmt"
)

// ====================
// Msg
// ====================
type Msg struct {
	Device  string  `json:"device"`
	Payload Payload `json:"payload"`
}

type Payload struct {
	Type    string      `json:"type"`
	Command string      `json:"command"`
	Result  bool        `json:"result,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
// ====================
// Device
// ====================
// Device is replica of Client{} object for create connection to device
type Device struct {
	Conn *websocket.Conn
	Send chan Msg
}

type Actioner interface {
	Action(Msg)
}

func (d Device) Write() {
	defer func() {
		d.Conn.Close()
	}()

	for {
		select {
		case msg, ok := <-d.Send:
			if !ok {
				d.Conn.WriteJSON(gin.H{"message": "Cannot Send message"})
				return
			}
			d.Conn.WriteJSON(msg)
		}
	}
}

func (d Device) Read() {
	msg := Msg{}
	for {
		err := d.Conn.ReadJSON(&msg)
		fmt.Println("Device command received: ", msg.Payload.Command)

		if err != nil {
			log.Println("Device Read JSON Error: ", msg)
			d.Conn.WriteJSON(gin.H{"message": "Read JSON Error: "})
			break
		}

		switch msg.Device {
		case "coin_hopper":
			h := CoinHopper{}
			h.Action(msg)
		case "coin_acceptor":
			c := CoinAcceptor{}
			c.Action(msg)
		case "bill_acceptor":
		case "printer":
		}

	}
}

type Acceptor interface {
	Serial() string
	Status() string
	CashReceive() int64
}
