package model

import (
	"fmt"
	"log"
)

type CoinAcceptor struct {
	machineId string
	Status    string
	Send      chan *Message
}

// Event & Response from coin acceptor.
func (ca *CoinAcceptor) event(s *Socket) {
	switch s.Msg.Command {
	case "received": // Event น้ีจะเกิดขึ้นเมื่อเคร่ืองรับเหรียญได้รับเหรียญ
		ca.Received(s)
	case "set_inhibit", "machine_id", "inhibit", "recently_inserted": // ตั้งค่า Inhibit (รับ-ไม่รับเหรียญ) ของ Coins Acceptor
		ca.Send <- s.Msg
	default:
		// "machine_id": 		// ร้องขอหมายเลข Serial Number ของ อุปกรณ์ Coins Acceptor
		// "inhibit":           // ร้องขอ สถานะ Inhibit (รับ-ไม่รับเหรียญ) ของ Coins Acceptor
		// "recently_inserted": // ร้องขอจานวนเงินของเหรียญล่าสุดที่ได้รับ
		ca.Send <- s.Msg
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

	m = <-ca.Send
	if !m.Result {
		log.Println("Error: coin Acceptor cannot start.")
		m.Command = "warning"
		m.Data = "Error: coin Acceptor cannot start."
		H.Web.Send <- m
		return
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

	m = <-ca.Send
	if !m.Result {
		log.Println("Error: coin Acceptor cannot stop.")
		m.Command = "warning"
		m.Data = "Error: coin Acceptor cannot stop."
		H.Web.Send <- m
		return
	}
	ca.Status = "STOP"
	fmt.Println("2. ปิดรับเหรียญสำเร็จ, CA status:", ca.Status)
}

func (ca *CoinAcceptor) Received(s *Socket) {
	fmt.Println("Start method: ca.Received() s.Msg.Data=", s.Msg.Data)
	//value := s.Msg.Data.(float64)
	value := ca.checkId(s.Msg.Data.(float64))
	PM.coin += value
	PM.total += value
	PM.remain -= value
	CB.hopper += value
	CB.total += value
	fmt.Println("PM.coin =", PM.coin, "PM total=", PM.total)
	PM.coinCh <- value

}

// checkId ตรวจเทียบ Id ที่ได้รับจากเครืื่องรับเหรียญ เทียบกับข้อมูลของผู้ผลิตส่งยอดเงินรับกลับ ปัจจุบันใช้เครื่อง MicroCoin SP115
func (ca *CoinAcceptor) checkId(data float64) float64 {
	var value float64
	switch data {
	case 1:
		value = 0.0
	case 2:
		value = 0.0
	case 3:
		value = 0.0
	case 4:
		value = 0.0
	case 5:
		value = 0.25
	case 6:
		value = 0.50
	case 7:
		value = 0.50
	case 8:
		value = 1.0
	case 9:
		value = 1.0
	case 10:
		value = 0.0
	case 11:
		value = 0.0
	case 12:
		value = 2.0
	case 13:
		value = 5.0
	case 14:
		value = 5.0
	case 15:
		value = 10.0
	case 16:
		value = 0.0
		// มี 1-16
	}
	return value
}