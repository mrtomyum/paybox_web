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
	fmt.Println("BillAcceptor Event...with Client=", c.Name)
	switch c.Msg.Command {
	case "machine_id": // ใช้สาหรับการร้องขอหมายเลข Serial Number ของ อุปกรณ์ Bill Acceptor
		switch c.Msg.Type {
		case "request":
			b.MachineId(H.Dev)
		case "response":
		}
	case "inhibit":           // ใช้สาหรับร้องขอ สถานะ Inhibit (รับ-ไม่รับธนบัตร) ของ Bill Acceptor
	case "set_inhibit":       // ตั้งค่า Inhibit (รับ-ไม่รับธนบัตร) ของ Bill Acceptor
	case "recently_inserted": // ร้องขอจานวนเงินของธนบัตรล่าสุดที่ได้รับ
	case "take_reject": // สั่งให้ รับ-คืน ธนบัตรท่ีกาลังพักอยู่
		BA.Send <- c.Msg
	case "received": // Event  นี้จะเกิดขึ้นเม่ือเคร่ืองรับธนบัตรได้รับธนบัตร
		OH.Bill = c.Msg.Data.(float64)
		OH.Total = + OH.Bill
		m := &Message{
			Command:"onhand",
			Data:   100,
		}
		H.Web.Send <- m
		fmt.Println("Bill Update")
	}
}

func (b *BillAcceptor) MachineId(c *Client) error {
	ch := make(chan *Message)
	m := &Message{Device:"bill_acc", Command:"machine_id", Type: "request"}
	c.Send <- m
	go func() {
		for {
			m = <-b.Send
			ch <- m
		}
	}()
	m = <-ch
	if !m.Result {
		b.Status = "Error when get machine_id"
		return errors.New("Error when get machine_id")
		log.Println("Error when get machine_id")
	}
	fmt.Println("Bill Acceptor machine id =", m.Data.(string))
	m.Type = "response"
	H.Web.Send <- m
	return nil
}

// สั่งให้ Bill Acceptor เก็บเงิน
func (b *BillAcceptor) Take(action bool) error {
	ch := make(chan *Message)
	m := &Message{
		Device:  "bill_acc",
		Command: "take_reject",
		Type:    "request",
		Data:    action,
	}
	H.Dev.Send <- m

	go func() {
		m = <-b.Send
		fmt.Println("1. Response from Bill Acceptor:")
		ch <- m
		fmt.Println("2...")
	}()

	fmt.Println("3. [*BillAcceptor.Take()] ส่ง m1-> c.Send รอคำตอบจาก Bill Acceptor, Message=", m)
	m = <-ch //  ที่นี่โปรแกรมจะ Block รอจนกว่าจะมี Message m3 จาก Channel ch
	close(ch)
	fmt.Println("4.... ")
	if !m.Result {
		b.Status = "Error cannot take bill"
		return errors.New("Error Bill Acceptor cannot take bill")
		log.Println("Error response from Bill Acceptor!")
	}
	fmt.Println("*BillAcceptor.Take() success.. m=:", m)
	return nil
}
