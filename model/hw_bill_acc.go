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
func (ba *BillAcceptor) event(s *Socket) {
	//fmt.Println("BillAcceptor Event...with Socket=", s.Name)
	switch s.Msg.Command {
	case "received": // Event  นี้จะเกิดขึ้นเม่ือเคร่ืองรับธนบัตรได้รับธนบัตร
		ba.Received(s)
	case "time_out":
		ba.TimeOut(s)
	case "set_inhibit", "machine_id", "inhibit", "recently_inserted", "take_reject": // ตั้งค่า Inhibit (รับ-ไม่รับธนบัตร) ของ bill Acceptor
		ba.Send <- s.Msg
	default:
		// "machine_id": 		// ใช้สาหรับการร้องขอหมายเลข Serial Number ของ อุปกรณ์ bill Acceptor
		// "inhibit":           // ใช้สาหรับร้องขอ สถานะ Inhibit (รับ-ไม่รับธนบัตร) ของ bill Acceptor
		// "recently_inserted": // ร้องขอจานวนเงินของธนบัตรล่าสุดที่ได้รับ
		// "take_reject": 		// สั่งให้ รับ-คืน ธนบัตรท่ีกาลังพักอยู่
		fmt.Println("BA Event() default:")
	}
}

// ใช้สาหรับการร้องขอหมายเลข Serial Number ของ อุปกรณ์ bill Acceptor
func (ba *BillAcceptor) MachineId(c *Socket) error {
	m := &Message{Device: "bill_acc", Command: "machine_id", Type: "request"}
	c.Send <- m
	m = <-ba.Send
	if !m.Result {
		ba.Status = "Error when get machine_id"
		log.Println("Error when get machine_id")
		return errors.New("Error when get machine_id")
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
		Data:    false,
	}
	fmt.Println("1...สั่งเปิดรับธนบัตรรอ response จาก BA")
	H.Hw.Send <- m
	// I/O blocking รอ HW ตอบกลับ
	m2 := <-ba.Send
	if !m2.Result {
		log.Println("Error: bill Acceptor cannot start. message =", m)
		m2.Command = "warning"
		m2.Data = "Error: bill Acceptor cannot start."
		H.Web.Send <- m2
		return
	}
	ba.Inhibit = false
	ba.Status = "START"
	fmt.Println("2. เปิดรับธนบัตรสำเร็จ, BA status:", ba.Status)
}

func (ba *BillAcceptor) Stop() {
	//ch := make(chan *Message)
	m := &Message{
		Device:  "bill_acc",
		Command: "set_inhibit",
		Type:    "request",
		Data:    true,
	}
	H.Hw.Send <- m
	//fmt.Println("1. สั่งปิดรับธนบัตรรอ response จาก BA...")
	// I/O blocking รอ HW ตอบกลับ
	m2 := <-ba.Send
	if !m2.Result {
		log.Println("Error: bill Acceptor cannot stop.")
		m2.Command = "warning"
		m2.Data = "Error: bill Acceptor cannot stop."
		H.Web.Send <- m2
		return
	}
	//ba.Inhibit = false
	//ba.Status = "STOP"
	//fmt.Println("2. ปิดรับธนบัตรสำเร็จ, BA status:", ba.Status)
}

var ErrNoBillEscrow error = errors.New("Error no bill escrowed = ไม่มีธนบัตรพัก")

func (ba *BillAcceptor) Take() {
	m := &Message{
		Device:  "bill_acc",
		Command: "take_reject",
		Type:    "request",
		Data:    true,
	}
	H.Hw.Send <- m
	//fmt.Printf("pm.Take() action = [%v] 1. รอคำตอบจาก bill Acceptor\n", action)
	//m = <-ba.Send
	//if !m.Result {
	//	ba.Status = "Error cannot take bill"
	//	log.Println("Error response from bill Acceptor!")
	//	return errors.New("Error bill Acceptor cannot take bill")
	//}
	//fmt.Printf("BA.Take() SUCCSS msg %v m.Result= %v, m.Data= %v\n", m, m.Result, m.Data)
}

func (ba *BillAcceptor) Reject() {
	m := &Message{
		Device:  "bill_acc",
		Command: "take_reject",
		Type:    "request",
		Data:    false,
	}
	H.Hw.Send <- m
	//fmt.Printf("ba.reject() 1. รอคำตอบจาก bill Acceptor\n")
	//	m = <-BA.Send
	//	if !m.Result {
	//		BA.Status = "Error cannot take bill"
	//		log.Println("Error response from bill Acceptor!")
	//		return errors.New("Error bill Acceptor cannot take bill")
	//	}
}

func (ba *BillAcceptor) Received(s *Socket) {
	fmt.Println("Start method: ba.send()")
	// todo: ตรวจ AcceptedBill ถ้า false ให้ BA.Reject()
	PM.send <- s.Msg
}

func (ba *BillAcceptor) TimeOut(s *Socket) {
	log.Println("Bill Acceptor -> Time Out")
	// Todo: send msg to UI to warning User
	go PM.Cancel(s)
}
