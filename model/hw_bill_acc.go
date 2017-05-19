package model

import (
	"errors"
	"fmt"
	"log"
)

type BillAcceptor struct {
	machineId string
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
	case "returned":
		ba.Returned(s)
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
func (ba *BillAcceptor) MachineId(s *Socket) error {
	m := &Message{Device: "bill_acc", Command: "machine_id", Type: "request"}
	s.Send <- m
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
		m2.Device = "web"
		m2.Command = "warning"
		m2.Data = "Error: bill Acceptor cannot stop."
		H.Web.Send <- m2
		return
	}
	ba.Status = "STOP"
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
	m = <-ba.Send
	if !m.Result {
		ba.Status = "Error cannot take bill"
		log.Println("Error response from bill Acceptor!")
		//return errors.New("Error bill Acceptor cannot take bill")
	}
	//fmt.Printf("BA.Take() SUCCSS msg %v m.Result= %v, m.Data= %v\n", m, m.Result, m.Data)
	// อัพเดตยอดเงินสดในตู้ด้วย
	CB.bill += PM.billEscrow  // เพิ่มยอดธนบัตรในถังธนบัตร
	CB.total += PM.billEscrow // เพิ่มยอดรวมของ CashBox
	PM.bill += PM.billEscrow
	PM.total += PM.billEscrow
	PM.remain -= PM.billEscrow
	PM.billEscrow = 0 // ล้างยอดเงินพัก
}

func (ba *BillAcceptor) Reject() {
	m := &Message{
		Device:  "bill_acc",
		Command: "take_reject",
		Type:    "request",
		Data:    false,
	}
	H.Hw.Send <- m
	fmt.Println("ba.reject() 1. รอคำตอบจาก bill Acceptor")
	m = <-BA.Send
	if !m.Result {
		BA.Status = "Error cannot take bill"
		log.Println("Error response from bill Acceptor!")
		//return errors.New("Error bill Acceptor cannot take bill")
	}
	// todo: Send msg to Ui to ask customer "กรุณาดึงธนบัตรออก" จนกว่าจะส่ง msg "returned" ไปให้ UI ให้ค้างคำเตือนกระพริบไว้
}

func (ba *BillAcceptor) Received(s *Socket) {
	fmt.Println("Start method: ba.Received()")
	value := ba.checkId(s.Msg.Data.(int))
	PM.billCh <- value
}

func (ba *BillAcceptor) TimeOut(s *Socket) {
	log.Println("Bill Acceptor -> Time Out")
	// Todo: billCh msg to UI to warning User
	go PM.Cancel() // ปิดไว้ก่อนมีบักจาก HW
	m := &Message{
		Command: "alert",
		Type:    "event",
		Data:    "คืนเงินเนื่องจากเป็นธนบัตรรุ่นใหม่ที่ไม่รู้จัก กรุณาเปลี่ยนธนบัตรใหม่ค่ะ",
	}
	H.Web.Send <- m
	log.Println("Send Notified Message to Web Ui: ", s.Msg)
}

func (ba *BillAcceptor) Returned(s *Socket) {
	// Send Message to Web Ui to notified
	m := &Message{
		Command: "alert",
		Type:    "event",
		Data:    "ลูกค้ารับธนบัตรคืนไปแล้ว",
	}
	H.Web.Send <- m
	log.Println("Send Notified Message to Web Ui: ", s.Msg)
	// todo: Log การคืนธนบัตรด้วย
}

// checkId ตรวจเทียบ Id ที่ได้รับจากเครืื่องรับธนบัตรเทียบกับข้อมูลของผู้ผลิตส่งยอดเงินกลับ
// ปัจจุบันใช้เครื่อง Global Bill Acceptors GBA ST2
func (ba *BillAcceptor) checkId(data int) float64 {
	var value float64
	switch data {
	case 1:
		value = 1.0
	case 2:
		value = 2.0
	case 3:
		value = 5.0
	case 4:
		value = 10.0
	case 5:
		value = 10.0
	case 6:
		value = 10.0
	case 7:
		value = 10.0
	case 8:
		value = 10.0
	case 9:
		value = 10.0
	case 10:
		value = 10.0
	case 11:
		value = 10.0
	case 12:
		value = 10.0
	case 13:
		value = 10.0
	case 14:
		value = 10.0
	case 15:
		value = 10.0
	case 16:
		value = 10.0
		// มี 1-16
	}
	return value
}