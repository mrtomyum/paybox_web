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
	case "set_inhibit", "machine_id", "inhibit", "recently_inserted", "take_reject": // ตั้งค่า Inhibit (รับ-ไม่รับธนบัตร) ของ Bill Acceptor
		ba.Send <- c.Msg
	default:
		// "machine_id": 		// ใช้สาหรับการร้องขอหมายเลข Serial Number ของ อุปกรณ์ Bill Acceptor
		// "inhibit":           // ใช้สาหรับร้องขอ สถานะ Inhibit (รับ-ไม่รับธนบัตร) ของ Bill Acceptor
		// "recently_inserted": // ร้องขอจานวนเงินของธนบัตรล่าสุดที่ได้รับ
		// "take_reject": 		// สั่งให้ รับ-คืน ธนบัตรท่ีกาลังพักอยู่
		fmt.Println("BA Event() default:")
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
	ba.machineId = m.Data.(string)
	m.Type = "response"
	H.Web.Send <- m
	return nil
}

func (ba *BillAcceptor) Start() {
	ch := make(chan *Message)
	m := &Message{
		Device:  "bill_acc",
		Command: "set_inhibit",
		Type:    "request",
		Data:    true,
	}
	fmt.Println("1...สั่งเปิดรับธนบัตรรอ response จาก BA")
	H.Dev.Send <- m
	go func() {
		m2 := <-ba.Send
		if !m2.Result {
			m2.Command = "warning"
			m2.Data = "Error: Bill Acceptor cannot start."
			H.Web.Send <- m2
			log.Println("Error: Bill Acceptor cannot start.")
		}
		ch <- m2
		return
	}()

	m = <-ch
	close(ch)
	ba.Inhibit = true
	ba.Status = "inhibit==true"
	fmt.Println("2. เปิดรับธนบัตรสำเร็จ, BA status:", ba.Status)
}

func (ba *BillAcceptor) Stop() {
	ch := make(chan *Message)
	m := &Message{
		Device:  "bill_acc",
		Command: "set_inhibit",
		Type:    "request",
		Data:    false,
	}
	H.Dev.Send <- m
	fmt.Println("1. สั่งปิดรับธนบัตรรอ response จาก BA...")
	go func() {
		m2 := <-ba.Send
		if !m2.Result {
			m2.Command = "warning"
			m2.Data = "Error: Bill Acceptor cannot stop."
			H.Web.Send <- m2
			log.Println("Error: Bill Acceptor cannot stop.")
		}
		ch <- m2
		return
	}()
	m = <-ch
	close(ch)
	ba.Inhibit = false
	ba.Status = "inhibit==false"
	fmt.Println("2. ปิดรับธนบัตรสำเร็จ, BA status:", ba.Status)
}

// สั่งให้ Bill Acceptor เก็บเงิน
func (ba *BillAcceptor) Take(action bool) error {
	ch := make(chan *Message)
	m1 := &Message{
		Device:  "bill_acc",
		Command: "take_reject",
		Type:    "request",
		Data:    action,
	}
	H.Dev.Send <- m1
	fmt.Printf("BA.Take() action = [%v] 1. รอคำตอบจาก Bill Acceptor", action)

	go func() {
		m2 := <-ba.Send
		fmt.Println("2. Response from Bill Acceptor:")
		ch <- m2
		//fmt.Println("3...")
	}()

	m3 := <-ch //  ที่นี่โปรแกรมจะ Block รอจนกว่าจะมี Message m3 จาก Channel ch
	close(ch)
	//fmt.Println("4. m3=", m3)
	if !m3.Result {
		ba.Status = "Error cannot take bill"
		return errors.New("Error Bill Acceptor cannot take bill")
		log.Println("Error response from Bill Acceptor!")
	}

	// แจ้ง Message -> Web
	//message := ""
	//switch action {
	//case true:
	//	message = "TAKE"
	//case false:
	//	message = "REJECT"
	//}
	//H.Web.Send <- &Message{
	//	Command: "warning",
	//	Data:    message,
	//}

	// อัพเดตยอดเงินสดในตู้ด้วย
	if m1.Data.(bool) == true { // ถ้าสั่ง Take
		CB.Bill += PM.BillEscrow  // เพิ่มยอดธนบัตรในถังธนบัตร
		CB.Total += PM.BillEscrow // เพิ่มยอดรวมของ CashBox
		PM.BillEscrow = 0         // ล้างยอดเงินพัก
	} else {
		PM.Total -= PM.BillEscrow // ลดยอดรับเงินรวม
		PM.Bill -= PM.BillEscrow  // ลดยอดรับธนบัตร
		PM.BillEscrow = 0         // ล้างยอดเงินพัก
	}

	fmt.Println("BA.Take() SUCCSS m1 =:", m1)
	return nil
}

func (ba *BillAcceptor) Received(c *Client) {
	fmt.Println("Start method: ba.Received()")
	received := c.Msg.Data.(float64)

	// todo: ตรวจ AcceptedBill ถ้า false ให้ BA.Reject()

	PM.Bill += received
	PM.BillEscrow = received
	PM.Total += received
	//m := &Message{
	//	Device:  "bill_acc",
	//	Command: "received",
	//	Data:    received,
	//}
	fmt.Printf("Sale = %v, Bill Received = %v, Bill Escrow=%v PM Total= %v\n", S.Total, PM.Bill, PM.BillEscrow, PM.Total)
	PM.Received <- c.Msg
	PM.OnHand(H.Web) // แจ้งยอดเงิน Payment กลับ Web
}
