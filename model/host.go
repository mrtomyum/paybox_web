package model

import (
	"errors"
	"fmt"
	"log"
)

type Host struct {
	Id            string
	Online        bool
	TotalEscrow   int // มูลค่าเงินพักทั้งหมด
	BillEscrow    int // มูลค่าธนบัตรที่พักอยู่ในเครื่องรับธนบัตร
	BillBox       int // มูลค่าธนบัตรในกล่องเก็บธนบัตร
	CoinHopperBox int // มูลค่าเหรียญใน Coin Hopper
	CoinBox       int // มูลค่าเหรียญใน CoinBox
	TotalCash     int // รวมมูลค่าเงินในตู้นี้
	Web           *Client
	Dev           *Client
	SetWebClient  chan *Client
	SetDevClient  chan *Client
}

func (h *Host) Run() {
	for {
		fmt.Println("Host.Run()...wait for next Channel")
		select {
		case c := <-h.SetWebClient:
			log.Println("<-Web:", c)
			h.Web = c
		case c := <-h.SetDevClient:
			log.Println("<-devClient:", c)
			h.Dev = c
		}
	}
}

// Onhand ส่งค่าเงินพัก Escrow ที่ Host เก็บไว้กลับไปให้ web
func (h *Host) Onhand(c *Client) {
	fmt.Println("Host.Onhand <-Message...")
	c.Msg.Result = true
	c.Msg.Type = "response"
	c.Msg.Data = h.TotalEscrow
	c.Send <- c.Msg
}

// Cancel คืนเงินจากทุก Device โดยตรวจสอบเงิน Escrow ใน Bill Acceptor ด้วยถ้ามีให้คืนเงิน
func (h *Host) Cancel(c *Client) error {
	fmt.Println("Host.Cancel()...")

	// Check Bill Acceptor
	if h.TotalEscrow == 0 { // ไม่มีเงินพัก
		log.Println("ไม่มีเงินพัก:")
		c.Msg.Type = "response"
		c.Msg.Result = false
		c.Msg.Data = "ไม่มีเงินพัก"
		c.Send <- c.Msg
		return errors.New("ไม่มีเงินพัก")
	}
	//  สั่งให้ BillAcceptor คืนเงินที่พักไว้
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
	coinHopperEscrow := h.TotalEscrow - h.BillEscrow
	h.BillEscrow = 0

	// CoinHopper
	// ให้จ่ายเหรียญที่คงค้างตามยอด coinHopperEscrow ออกด้านหน้า
	m2 := &Message{
		Device:  "coin_hopper",
		Command: "payout_by_cash",
		Type:    "request",
		Data:    coinHopperEscrow,
	}
	h.Dev.Send <- m2

	// Todo: Check if error from CoinHopper
	err = h.Dev.Ws.ReadJSON(&m2)
	if err != nil {
		log.Println("Cancel() Coin Hopper error:", err)
		c.Msg.Result = false
		c.Msg.Type = "response"
		c.Msg.Data = m2.Data
		c.Send <- c.Msg
		return err
	}
	h.TotalEscrow = 0 // เคลียร์ยอดเงินค้างให้หมด

	// Send message response back to Web Client
	c.Msg.Type = "response"
	c.Msg.Result = true
	c.Msg.Data = "sucess"
	c.Send <- c.Msg
	return nil
}
