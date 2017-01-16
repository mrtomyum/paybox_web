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
	m := &Message{
		Device:  "bill_acceptor",
		Command: "take_reject",
		Type:    "request",
		Data:    true,
	}
	c.Send <- m

	go func() {
		for {
			select {
			case m = <-b.Send:
				fmt.Println("Received response from Bill Acceptor:")
				if !m.Result {
					b.Status = "Error cannot take bill"
					log.Println("Error response from Bill Acceptor!")
				}
				H.TotalBill = + H.BillEscrow
				H.BillEscrow = 0
				fmt.Println("Bill Acc [take] success...Received response from Bill Acceptor:", m.Result)
				break
			}
		}
	}()
}
