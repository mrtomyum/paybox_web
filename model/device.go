package model

import (
	"log"
	"github.com/gorilla/websocket"
	"github.com/gin-gonic/gin"
	"fmt"
)

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
	m := Msg{}
	for {
		err := d.Conn.ReadJSON(&m)
		fmt.Println("Device command received: ", m.Payload.Command)

		if err != nil {
			log.Println("Device Read JSON Error: ", m)
			d.Conn.WriteJSON(gin.H{"message": "Read JSON Error: "})
			break
		}

		switch m.Device {
		case "coin_hopper":
			h := CoinHopper{}
			h.Action(d, m)
		case "coin_acceptor":
			ca := CoinAcceptor{}
			ca.Action(d, m)
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
