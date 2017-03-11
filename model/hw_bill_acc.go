package model

import (
	"log"
	"fmt"
	"errors"
)

type BillAcceptor struct {
	machineId string `json:"machine_id"`
	Inhibit   bool
	Status    string
	Send      chan *Message
}

// Event & Response from bill acceptor.
// อุปกรณ์จะทำงานครั้งละ 1 command อยู่แล้ว
// ดังนั้นไม่ต้องกลัวจะมี Event หรือ Response ข้ามลำดับกัน
func (ba *BillAcceptor) Event(c *Client) {
	fmt.Println("BillAcceptor Event...with Client=", c.Name)
	switch c.Msg.Command {
	case "received": // Event  นี้จะเกิดขึ้นเม่ือเคร่ืองรับธนบัตรได้รับธนบัตร
		ba.Received(c)
	case "time_out":
		ba.TimeOut(c)
	case "set_inhibit", "machine_id", "inhibit", "recently_inserted", "take_reject": // ตั้งค่า Inhibit (รับ-ไม่รับธนบัตร) ของ bill Acceptor
		ba.Send <- c.Msg
	default:
		// "machine_id": 		// ใช้สาหรับการร้องขอหมายเลข Serial Number ของ อุปกรณ์ bill Acceptor
		// "inhibit":           // ใช้สาหรับร้องขอ สถานะ Inhibit (รับ-ไม่รับธนบัตร) ของ bill Acceptor
		// "recently_inserted": // ร้องขอจานวนเงินของธนบัตรล่าสุดที่ได้รับ
		// "take_reject": 		// สั่งให้ รับ-คืน ธนบัตรท่ีกาลังพักอยู่
		fmt.Println("BA Event() default:")
	}
}

// ใช้สาหรับการร้องขอหมายเลข Serial Number ของ อุปกรณ์ bill Acceptor
func (ba *BillAcceptor) MachineId(c *Client) error {
	m := &Message{Device: "bill_acc", Command: "machine_id", Type: "request"}
	c.Send <- m
	//go func() {
	m = <-ba.Send
	//ch <- m
	//}()
	//m = <-ch
	if !m.Result {
		ba.Status = "Error when get machine_id"
		return errors.New("Error when get machine_id")
		log.Println("Error when get machine_id")
	}
	fmt.Println("bill Acceptor machine id =", m.Data.(string))
	ba.machineId = m.Data.(string)
	m.Type = "response"
	H.Web.Send <- m
	return nil
}

func (ba *BillAcceptor) Start() {
	//ch := make(chan *Message)
	m := &Message{
		Device:  "bill_acc",
		Command: "set_inhibit",
		Type:    "request",
		Data:    true,
	}
	fmt.Println("1...สั่งเปิดรับธนบัตรรอ response จาก BA")
	H.Hw.Send <- m
	//go func() {
	m2 := <-ba.Send
	if !m2.Result {
		m2.Command = "warning"
		m2.Data = "Error: bill Acceptor cannot start."
		H.Web.Send <- m2
		log.Println("Error: bill Acceptor cannot start.")
	}
	//	}
	//	ch <- m2
	//	return
	//}()
	//
	//m = <-ch
	//close(ch)
	ba.Inhibit = true
	ba.Status = "START"
	fmt.Println("2. เปิดรับธนบัตรสำเร็จ, BA status:", ba.Status)
}

func (ba *BillAcceptor) Stop() {
	//ch := make(chan *Message)
	m := &Message{
		Device:  "bill_acc",
		Command: "set_inhibit",
		Type:    "request",
		Data:    false,
	}
	H.Hw.Send <- m
	fmt.Println("1. สั่งปิดรับธนบัตรรอ response จาก BA...")
	//go func () {
	m2 := <-ba.Send
	if !m2.Result {
		m2.Command = "warning"
		m2.Data = "Error: bill Acceptor cannot stop."
		H.Web.Send <- m2
		log.Println("Error: bill Acceptor cannot stop.")
	}
	//	ch <- m2
	//	return
	//}()
	//m = <-ch
	//close(ch)
	ba.Inhibit = false
	ba.Status = "STOP"
	fmt.Println("2. ปิดรับธนบัตรสำเร็จ, BA status:", ba.Status)
}

var ErrNoBillEscrow error = errors.New("Error no bill escrowed = ไม่มีธนบัตรพัก")

// สั่งให้ bill Acceptor เก็บเงิน
func (ba *BillAcceptor) Take(action bool) error {
	if PM.billEscrow == 0 { // ถ้ามีธนบัตรพักอยู่ ให้สั่งเก็บธนบัตร
		return ErrNoBillEscrow
	}
	//ch := make(chan *Message)
	m := &Message{
		Device:  "bill_acc",
		Command: "take_reject",
		Type:    "request",
		Data:    action,
	}
	H.Hw.Send <- m
	fmt.Printf("BA.Take() action = [%v] 1. รอคำตอบจาก bill Acceptor", action)

	//go func() {
	m = <-ba.Send
	//fmt.Println("2. Response from bill Acceptor:")
	//ch <- m2
	//}()
	//
	//m3 := <-ch //  ที่นี่โปรแกรมจะ Block รอจนกว่าจะมี Message m3 จาก Channel ch
	//close(ch)
	if !m.Result {
		ba.Status = "Error cannot take bill"
		log.Println("Error response from bill Acceptor!")
		return errors.New("Error bill Acceptor cannot take bill")
	}

	// อัพเดตยอดเงินสดในตู้ด้วย
	//if m.Result { // ถ้าสั่ง Take
	CB.bill += PM.billEscrow  // เพิ่มยอดธนบัตรในถังธนบัตร
	CB.total += PM.billEscrow // เพิ่มยอดรวมของ CashBox
	PM.bill += PM.billEscrow
	PM.total += PM.billEscrow
	PM.remain -= PM.billEscrow
	PM.billEscrow = 0 // ล้างยอดเงินพัก
	//} else {
	//	PM.billEscrow = 0 // ล้างยอดเงินพัก
	//}

	fmt.Println("BA.Take() SUCCSS m1 =:", m)
	return nil
}

func (ba *BillAcceptor) Received(c *Client) {
	fmt.Println("Start method: ba.receivedCh()")
	received := c.Msg.Data.(float64)

	// todo: ตรวจ AcceptedBill ถ้า false ให้ BA.Reject()

	PM.billEscrow = received
	fmt.Printf("Sale = %v, bill receivedCh = %v, bill Escrow=%v PM total= %v\n", S.Total, PM.bill, PM.billEscrow, PM.total)
	PM.receivedCh <- c.Msg
	PM.OnHand(H.Web) // แจ้งยอดเงิน Payment กลับ Web
}

func (ba *BillAcceptor) TimeOut(c *Client) {
	PM.Cancel(c)
}
