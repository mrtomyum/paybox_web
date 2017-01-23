package model

import (
	"errors"
	"fmt"
	"log"
)

type Host struct {
	Id               string  // รหัสเมนบอร์ดตู้
	IsNetOnline      bool    // สถานะ GSM ปัจจุบัน (Real time)
	IsServerOnline   bool    // สถานะเซิร์ฟเวอร์ครั้งสุดท้ายที่สื่อสาร
	Web              *Client // Web Client object ที่เปิดคอนเนคชั่นอยู่
	Dev              *Client // Device Client object ที่เปิดคอนเนคชั่นอยู่
	LastTicketNumber int     // เลขคิวตั๋วล่าสุด ของแต่ละวัน ปิดเครื่องต้องยังอยู่ ขึ้นวันใหม่ต้อง Reset
}

// TotalEscrow ส่งค่าเงินพัก Escrow ที่ Host เก็บไว้กลับไปให้ web
func (h *Host) OnHand(web *Client) {
	fmt.Println("Host.OnHand()...")
	web.Msg.Command = "onhand"
	web.Msg.Result = true
	web.Msg.Type = "event"
	web.Msg.Data = PM.Total
	web.Send <- web.Msg
}

// Cancel คืนเงินจากทุก Device โดยตรวจสอบเงิน Escrow ใน Bill Acceptor ด้วยถ้ามีให้คืนเงิน
func (h *Host) Cancel(c *Client) error {
	fmt.Println("Host.Cancel()...")

	// Check Bill Acceptor
	if PM.Total == 0 { // ไม่มีเงินพัก
		log.Println("ไม่มีเงินพัก:")
		c.Msg.Type = "response"
		c.Msg.Result = false
		c.Msg.Data = "ไม่มีเงินพัก"
		c.Send <- c.Msg
		return errors.New("ไม่มีเงินพัก")
	}
	// สั่งให้ BillAcceptor คืนเงินที่พักไว้
	m1 := &Message{
		Device:  "bill_acc",
		Command: "escrow",
		Type:    "request",
		Result:  true,
		Data:    false,
	}
	h.Dev.Send <- m1

	// Check BillAcc response
	err := h.Dev.Ws.ReadJSON(&m1)
	if err != nil {
		log.Println("Host.Cancel() error ->", m1.Data)
		return err
	}

	// Success
	PM.Coin = PM.Total - PM.Bill
	PM.Bill = 0

	// CoinHopper สั่งให้จ่ายเหรียญที่คงค้างตามยอด coinHopperEscrow ออกด้านหน้า
	m2 := &Message{
		Device:  "coin_hopper",
		Command: "payout_by_cash",
		Type:    "request",
		Data:    PM.Coin,
	}
	h.Dev.Send <- m2

	// Check if error from CoinHopper
	err = h.Dev.Ws.ReadJSON(&m2)
	if err != nil {
		log.Println("Cancel() Coin Hopper error:", err)
		c.Msg.Result = false
		c.Msg.Type = "response"
		c.Msg.Data = m2.Data
		c.Send <- c.Msg
		return err
	}
	PM.Total = 0 // เคลียร์ยอดเงินค้างให้หมด

	// Send message response back to Web Client
	c.Msg.Type = "response"
	c.Msg.Result = true
	c.Msg.Data = "sucess"
	c.Send <- c.Msg
	return nil
}

