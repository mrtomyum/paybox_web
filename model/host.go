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
	coinHopperEscrow := h.TotalEscrow - h.BillEscrow
	h.BillEscrow = 0

	// CoinHopper สั่งให้จ่ายเหรียญที่คงค้างตามยอด coinHopperEscrow ออกด้านหน้า
	m2 := &Message{
		Device:  "coin_hopper",
		Command: "payout_by_cash",
		Type:    "request",
		Data:    coinHopperEscrow,
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
	h.TotalEscrow = 0 // เคลียร์ยอดเงินค้างให้หมด

	// Send message response back to Web Client
	c.Msg.Type = "response"
	c.Msg.Result = true
	c.Msg.Data = "sucess"
	c.Send <- c.Msg
	return nil
}

// Billing ทำการบันทึกรับชำระเงิน โดยตรวจสอบการ ทอนเงิน บันทึกลง SqLite
// และส่งข้อมูล Order Post ขึ้น Cloud แต่หาก Network Down Order.completed = false
// จะมี Routine Check Network status  คอยตรวจสอบสถานะและ Retry
func (h *Host) Billing(c *Client) error {
	// สั่ง device ทอนเงิน
	// บันทึกข้อมูลลง SQL โดย order.completed = false
	// Post Order ขึ้น Cloud
	// Check Network Status
	// ถ้า Net down หรือ Post สำเร็จ ให้บันทึก SQL order.completed = true
	return nil
}