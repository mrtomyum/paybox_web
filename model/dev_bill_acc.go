package model

import (
	"log"
	"fmt"
	"errors"
)

type BillAcceptor struct {
	Id     string
	Status string
	Send   chan *Message
}

func (b *BillAcceptor) Event(c *Client) {
	switch c.Msg.Command {
	case "machine_id":        // ใช้สาหรับการร้องขอหมายเลข Serial Number ของ อุปกรณ์ Bill Acceptor
	case "inhibit":           // ใช้สาหรับร้องขอ สถานะ Inhibit (รับ-ไม่รับธนบัตร) ของ Bill Acceptor
	case "set_inhibit":       // ตั้งค่า Inhibit (รับ-ไม่รับธนบัตร) ของ Bill Acceptor
	case "recently_inserted": // ร้องขอจานวนเงินของธนบัตรล่าสุดที่ได้รับ
	case "take_reject": // สั่งให้ รับ-คืน ธนบัตรท่ีกาลังตรวจสอบอยู่ **น่าจะใช้คำว่า Escrow
		B.Send <- c.Msg
	case "received": // Event นจี้ ะเกิดขึ้นเม่ือเคร่ืองรับธนบัตรได้รับธนบัตร
		H.BillEscrow = c.Msg.Data.(float64)
		m := &Message{
			Command:"onhand",
			Data:   100,
		}
		H.Web.Send <- m
		fmt.Println("Bill Update")
	}
}

// สั่งให้ Bill Acceptor เก็บเงิน
func (b *BillAcceptor) Take(c *Client) error {
	ch := make(chan *Message)
	m1 := &Message{
		Device:  "bill_acc",
		Command: "take_reject",
		Type:    "request",
		Data:    true,
	}
	fmt.Println("[*BillAcceptor.Take()] ส่ง m1-> c.Send รอคำตอบจาก Bill Acceptor, Message=", m1)
	c.Send <- m1

	go func() {
		for {
			select {
			case m2 := <-b.Send:
				fmt.Println("Received response from Bill Acceptor:")
				ch <- m2
				break
			}
		}
	}()
	m3 := <-ch //  ที่นี่โปรแกรมจะ Block รอจนกว่าจะมี Message m3 จาก Channel ch
	if !m3.Result {
		b.Status = "Error cannot take bill"
		return errors.New("Error Bill Acceptor cannot take bill")
		log.Println("Error response from Bill Acceptor!")
	}
	H.TotalBill = + H.BillEscrow
	H.BillEscrow = 0
	fmt.Println("Bill Acceptor [take] success.. m3=:", m3)
	return nil
}
