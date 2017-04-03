package model

import (
	"log"
	"fmt"
)

type CoinAcceptor struct {
	machineId string `json:"machine_id"`
	Inhibit   bool
	Status    string
	Send      chan *Message
}

// Event & Response from coin acceptor.
func (ca *CoinAcceptor) event(c *Socket) {
	switch c.Msg.Command {
	case "received": // Event น้ีจะเกิดขึ้นเมื่อเคร่ืองรับเหรียญได้รับเหรียญ
		ca.Received(c)
	case "set_inhibit", "machine_id", "inhibit", "recently_inserted": // ตั้งค่า Inhibit (รับ-ไม่รับเหรียญ) ของ Coins Acceptor
		ca.Send <- c.Msg
	default:
		// "machine_id": 		// ร้องขอหมายเลข Serial Number ของ อุปกรณ์ Coins Acceptor
		// "inhibit":           // ร้องขอ สถานะ Inhibit (รับ-ไม่รับเหรียญ) ของ Coins Acceptor
		// "recently_inserted": // ร้องขอจานวนเงินของเหรียญล่าสุดที่ได้รับ
		ca.Send <- c.Msg
	}
}

func (ca *CoinAcceptor) Start() {
	//ch := make(chan *Message)
	m := &Message{
		Device:  "coin_acc",
		Command: "set_inhibit",
		Type:    "request",
		Data:    false,
	}
	H.Hw.Send <- m
	fmt.Println("1. สั่งเปิดรับเหรียญรอ response จาก CA...")
	//go func() {
	m = <-ca.Send
	if !m.Result {
		log.Println("Error: coin Acceptor cannot start.")
		m.Command = "warning"
		m.Data = "Error: coin Acceptor cannot start."
		H.Web.Send <- m
	}
	ca.Status = "START"
	fmt.Println("2. เปิดรับเหรียญสำเร็จ, CA status:", ca.Status)
}

func (ca *CoinAcceptor) Stop() {
	//ch := make(chan *Message)
	m := &Message{
		Device:  "coin_acc",
		Command: "set_inhibit",
		Type:    "request",
		Data:    true,
	}
	H.Hw.Send <- m
	fmt.Println("1. สั่งปิดรับเหรียญรอ response จาก CA...")
	//go func() {
	m = <-ca.Send
	if !m.Result {
		log.Println("Error: coin Acceptor cannot stop.")
		m.Command = "warning"
		m.Data = "Error: coin Acceptor cannot stop."
		H.Web.Send <- m
	}
	ca.Status = "STOP"
	fmt.Println("2. ปิดรับเหรียญสำเร็จ, CA status:", ca.Status)
}

func (ca *CoinAcceptor) Received(c *Socket) {
	fmt.Println("Start method: ca.receivedCh()")
	PM.receivedCh <- c.Msg
}
