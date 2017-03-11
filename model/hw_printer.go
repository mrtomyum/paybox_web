package model

import (
	"fmt"
	"errors"
	"time"
)

type Printer struct {
	machineId string `json:"machine_id"`
	Status    string
	Send      chan *Message
}

func (p *Printer) Event(c *Socket) {
	switch c.Msg.Command {
	case "machine_id", "do_single", "do_group":
		p.Send <- c.Msg
	//case "machine_id": // ร้องขอหมายเลข Serial Number ของ อุปกรณ์ Printer
	//case "do_single":  //ส่ังการเคร่ืองปริ้นเตอร์ แบบส่งคาส่ังการกระทาคาสั่งเดียว โดย action_name และ action_data สามารถดูได้จากตาราง Action
	//case "do_group":   //ส่ังการเคร่ืองปร้ินเตอร์ แบบส่งคาส่ังการกระทาแบบเปน็ ชุด โดย action_name และ action_data สามารถดูได้จากตาราง Action
	case "near_end": // Event แจ้งเตือนกระดาษใกล้หมด
	case "no_paper": // Event แจ้งเตือนกระดาษหมดแล้ว
	}
}

func (p *Printer) Print(s *Sale) error {
	fmt.Println("p.Print() run")
	data, err := p.makeSaleSlip(s)
	if err != nil {
		return err
	}

	m := &Message{
		Device:  "printer",
		Command: "do_group",
		Type:    "request",
		Data:    data,
	}
	H.Hw.Send <- m
	fmt.Println("1. สั่งพิมพ์ รอ Priner ตอบสนอง")
	//go func() {
	m = <-p.Send
	//ch <- m2
	//}()
	//m = <-ch
	if !m.Result {
		return errors.New("Err: printer error.")
	}
	fmt.Println("พิมพ์สำเร็จ Print success!")
	m2 := &Message{
		Device:  "web",
		Command: "print",
		Type:    "event",
		Data:    "success",
	}
	H.Web.Send <- m2
	return nil
}

func (p *Printer) PrintTest(data string) error {
	fmt.Println("p.PrintTest() run")

	// หน่วงเวลารอ Host เชื่อม Websocket -> Device ให้เสร็จก่อน
	timer := time.NewTimer(time.Millisecond * 100)
	<-timer.C

	m := &Message{
		Device:  "printer",
		Command: "do_single",
		Type:    "request",
		Data:    data,
	}
	H.Hw.Send <- m
	fmt.Println("1. สั่งพิมพ์ รอ Priner ตอบสนอง")

	m = <-p.Send
	if !m.Result {
		return errors.New("Err: printer error.")
	}
	fmt.Println("พิมพ์สำเร็จ Print success!")
	m2 := &Message{
		Device:  "web",
		Command: "print_test",
		Type:    "event",
		Data:    "success",
	}
	H.Web.Send <- m2
	return nil
}

func (p *Printer) makeSaleSlip(s *Sale) (data string, err error) {
	header := `[{"action":"set_text_size","action_data":3},{"action":"printline", "action_data":"ร้านกาแฟ MOMO"},{"action":"set_text_size","action_data":1},{"action":"printline", "action_data":"Ticketid  : 12"},{"action":"printline", "action_data": "ID     NAME      QTY     AMT"},`
	//header := `[{"set_text_size":3},{"printline" : "ร้านกาแฟ MOMO"},{"set_text_size":1},{"printline": "Ticketid  : 12"},{"printline": "ID     NAME      QTY     AMT"},`
	item := `{"action":"printline", "action_data": "2     Late        1       40.00"},`
	footer := `{"action":"printline", "action_data": "----------------------"},{"action":"printline", "action_data": "รวมมูลค่าสินค้า     %v"},{"action":"printline","action_data": "รับเงิน           %v"},{"action":"printline", "action_data": "เงินทอน          %v"},{"action":"printline", "action_data": "ขอบคุณที่ใช้บริการ"},{"action":"paper_cut","action_data": {"type": "partial_cut","feed": 1}},`
	queue := `{"action":"printline","action_data": "Ticket" },{"action":"set_text_size","action_data":8},{"action":"paper_cut","action_data": {"type": "full_cut","feed": 1}}]`
	//data = fmt.Sprintf(header+item+footer+queue, s.total, s.New, s.Change)
	//data = header
	data = header + item + footer + queue
	fmt.Println("data=", data)
	return data, nil
}

func (p *Printer) makeTicket(s *Sale) error {
	return nil
}

func (p *Printer) makeRefund(value float64) error {
	return nil
}
