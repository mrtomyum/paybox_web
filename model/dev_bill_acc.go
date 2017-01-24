package model

import (
	"log"
	"fmt"
	"errors"
)

type BillAcceptor struct {
	Id      string
	Inhibit bool
	Status  string
	Send    chan *Message
}

// Event & Response from bill acceptor.
// อุปกรณ์จะทำงานครั้งละ 1 command อยู่แล้ว
// ดังนั้นไม่ต้องกลัวจะมี Event หรือ Response ข้ามลำดับกัน
func (ba *BillAcceptor) Event(c *Client) {
	fmt.Println("BillAcceptor Event...with Client=", c.Name)
	switch c.Msg.Command {
	case "received": // Event  นี้จะเกิดขึ้นเม่ือเคร่ืองรับธนบัตรได้รับธนบัตร
		ba.Received(c)
	default:
		// "machine_id": 		// ใช้สาหรับการร้องขอหมายเลข Serial Number ของ อุปกรณ์ Bill Acceptor
		// "inhibit":           // ใช้สาหรับร้องขอ สถานะ Inhibit (รับ-ไม่รับธนบัตร) ของ Bill Acceptor
		// "set_inhibit":       // ตั้งค่า Inhibit (รับ-ไม่รับธนบัตร) ของ Bill Acceptor
		// "recently_inserted": // ร้องขอจานวนเงินของธนบัตรล่าสุดที่ได้รับ
		// "take_reject": 		// สั่งให้ รับ-คืน ธนบัตรท่ีกาลังพักอยู่
		ba.Send <- c.Msg
	}
}

// ใช้สาหรับการร้องขอหมายเลข Serial Number ของ อุปกรณ์ Bill Acceptor
func (ba *BillAcceptor) MachineId(c *Client) error {
	ch := make(chan *Message)
	m := &Message{Device:"bill_acc", Command:"machine_id", Type: "request"}
	c.Send <- m
	go func() {
		for {
			m = <-ba.Send
			ch <- m
		}
	}()
	m = <-ch
	if !m.Result {
		ba.Status = "Error when get machine_id"
		return errors.New("Error when get machine_id")
		log.Println("Error when get machine_id")
	}
	fmt.Println("Bill Acceptor machine id =", m.Data.(string))
	m.Type = "response"
	H.Web.Send <- m
	return nil
}

func (ba *BillAcceptor) Start() {
	ch := make(chan *Message)
	m := &Message{
		Device:  "bill_acc",
		Command: "set_inhibit",
		Type:    "response",
		Data:    true,
	}
	ba.Send <- m
	go func() {
		m2 := <-ba.Send
		if !m2.Result {
			m2.Command = "warning"
			m2.Data = "Error: Bill Acceptor cannot start."
			H.Web.Send <- m2
		}
		log.Println("Error: Bill Acceptor cannot start.")
		ch <- m2
		return
	}()
	m = <-ch
	close(ch)
	ba.Inhibit = true
}

func (ba *BillAcceptor) Stop() {
	ch := make(chan *Message)
	m := &Message{
		Device:  "bill_acc",
		Command: "set_inhibit",
		Data:    false,
	}
	ba.Send <- m
	go func() {
		m2 := <-ba.Send
		if !m2.Result {
			m2.Command = "warning"
			m2.Data = "Error: Bill Acceptor cannot stop."
			H.Web.Send <- m2
		}
		log.Println("Error: Bill Acceptor cannot stop.")
		ch <- m2
		return
	}()
	m = <-ch
	close(ch)
	ba.Inhibit = false
}

// สั่งให้ Bill Acceptor เก็บเงิน
func (ba *BillAcceptor) Take(action bool) error {
	ch := make(chan *Message)
	m := &Message{
		Device:  "bill_acc",
		Command: "take_reject",
		Type:    "request",
		Data:    action,
	}
	H.Dev.Send <- m

	go func() {
		m = <-ba.Send
		fmt.Println("1. Response from Bill Acceptor:")
		ch <- m
		fmt.Println("2...")
	}()

	fmt.Println("3. [*BillAcceptor.Take()] ส่ง m1-> c.Send รอคำตอบจาก Bill Acceptor, Message=", m)
	m = <-ch //  ที่นี่โปรแกรมจะ Block รอจนกว่าจะมี Message m3 จาก Channel ch
	close(ch)
	fmt.Println("4.... ")
	if !m.Result {
		ba.Status = "Error cannot take bill"
		return errors.New("Error Bill Acceptor cannot take bill")
		log.Println("Error response from Bill Acceptor!")
	}
	// อัพเดตยอดเงินสดในตู้ด้วย
	CB.Bill = + PM.Bill
	CB.Total = + PM.Bill
	PM.Total = - PM.Bill
	PM.Bill = 0
	fmt.Println("*BillAcceptor.Take() success.. m=:", m)
	return nil
}

func (ba *BillAcceptor) Received(c *Client) {
	fmt.Println("Start method: ba.Received()")
	received := c.Msg.Data.(float64)
	PM.Bill += received
	PM.Total += received
	//m := &Message{
	//	Device:  "bill_acc",
	//	Command: "received",
	//	Data:    received,
	//}
	fmt.Printf("Bill Received = %v, PM Total= %v\n", PM.Bill, PM.Total)
	//PM.Send <- m
	PM.OnHand(H.Web) // แจ้งยอดเงิน Payment กลับ Web
}
