package model

import (
	"log"
	"fmt"
)

type BillAcceptor struct {
	Id     string
	Status string
	Send   chan *Message
}

// สั่งให้ Bill Acceptor เก็บเงิน
func (b *BillAcceptor) Take(c *Client) {
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
		log.Println("Error response from Bill Acceptor!")
	}
	H.TotalBill = + H.BillEscrow
	H.BillEscrow = 0
	fmt.Println("Bill Acc [take] success...Received response from Bill Acceptor:", m3)
}
