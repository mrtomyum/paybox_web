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

//case "machine_id": // ร้องขอหมายเลข Serial Number ของ อุปกรณ์ Printer
//case "do_single":  //ส่ังการเคร่ืองปริ้นเตอร์ แบบส่งคาส่ังการกระทาคาสั่งเดียว โดย action_name และ action_data สามารถดูได้จากตาราง Action
//case "do_group":   //ส่ังการเคร่ืองปร้ินเตอร์ แบบส่งคาส่ังการกระทาแบบเปน็ ชุด โดย action_name และ action_data สามารถดูได้จากตาราง Action
func (p *Printer) Event(c *Socket) {
	switch c.Msg.Command {
	case "machine_id", "do_single", "do_group":
		p.Send <- c.Msg
	case "near_end": // Event แจ้งเตือนกระดาษใกล้หมด
		m := &Message{
			Device:  "web",
			Command: "near_end",
			Type:    "event",
			Data:    "Paper near end // คำเตือน กระดาษใกล้หมด โปรดแจ้งพนักงาน",
		}
		H.Web.Send <- m
	case "no_paper": // Event แจ้งเตือนกระดาษหมดแล้ว
		m := &Message{
			Device:  "web",
			Command: "no_paper",
			Type:    "event",
			Data:    "Paper run out// คำเตือน กระดาษหมด!! กรุณาแจ้งพนักงานให้เปลี่ยนกระดาษ",
		}
		H.Web.Send <- m
	}
}

func (p *Printer) makeTicket(s *Sale) doGroup {
	var g doGroup
	g.setTextSize(0)
	g.printLine("============ ร้านกาแฟ สุขใจขายได้สบายดี ============")
	var today = time.Now()
	date := fmt.Sprintf("%s", today.Format(time.RFC3339Nano))
	g.printLine(date)
	ss := s.SaleSubs
	for _, sub := range ss {
		item := fmt.Sprintf("%2dx%-17s%5s", sub.Qty, sub.ItemName, sub.PriceName)
		detail := fmt.Sprintf("@%3.2f฿ = %4.2f฿", sub.Price, float64(sub.Qty)*sub.Price)
		g.setTextSize(1)
		g.printLine(item)
		g.setTextSize(0)
		g.print(detail)
		g.newline()
	}
	sum := fmt.Sprintf("%6s%4.2f%6s%6.2f%6s%6.2f", "Total:", s.Total, "Payment:", s.Pay, "Change:", s.Change)
	g.print(sum)
	g.newline()
	g.printLine("================================================")
	//g.printBarcode("CODE39", "12345678")
	g.paperCut("full_cut", 90)
	fmt.Println(&g.actions)
	return g
}

// Print() รับค่า Sale แล้วพิมพ์ฟอร์มที่กำหนด ให้
func (p *Printer) PrintTicket(s *Sale) error {
	fmt.Println("p.Print() run")
	data := p.makeTicket(s)
	//data := gin.H{"action": "printline", "action_data": "นี่คือคูปอง"}
	m := &Message{
		Device:  "printer",
		Command: "do_group",
		Type:    "request",
		Data:    data.actions,
	}

	H.Hw.Send <- m
	fmt.Println("1. สั่งพิมพ์ รอ Priner ตอบสนอง...")
	m = <-p.Send
	if !m.Result {
		m2 := &Message{
			Device:  "web",
			Command: "print",
			Type:    "event",
			Data:    "Print error",
		}
		H.Web.Send <- m2
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

func (p *Printer) makeRefund(value float64) error {
	return nil
}

//=============== ACTION ====================
type action struct {
	Name string `json:"action"`
	Data interface{} `json:"action_data"`
}

//=============== DO_GROUP ====================
type doGroup struct {
	actions []*action
}

func (g *doGroup) print(s string) {
	a := &action{
		Name: "print",
		Data: s,
	}
	g.actions = append(g.actions, a)
}

func (g *doGroup) printLine(s string) {
	a := &action{"printline", s}
	g.actions = append(g.actions, a)
}

func (g *doGroup) setTextSize(size int) {
	a := &action{"set_text_size", size}
	g.actions = append(g.actions, a)
}

func (g *doGroup) newline() {
	a := &action{Name: "newline"}
	g.actions = append(g.actions, a)
}

//=========== BARCODE =============
type barcode struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

func (g *doGroup) printBarcode(t, d string) {
	data := barcode{t, d}
	a := &action{
		Name: "print_barcode",
		Data: data,
	}
	g.actions = append(g.actions, a)
}

//=========== QR_CODE =============
type qrCode struct {
	Mag      int `json:"mag"`
	Ecl      int `json:"ect"`
	DataType string `json:"data_type"`
	Data     string `json:"data"`
}

func (g *doGroup) printQr(mag, ecl int, data_type, data string) {
	d := qrCode{mag, ecl, data_type, data}
	a := &action{
		Name: "print_qr",
		Data: d,
	}
	g.actions = append(g.actions, a)
}

//=========== PAPER_CUT =============
type paperCut struct {
	Type string `json:"type"`
	Feed int `json:"feed"`
}

func (g *doGroup) paperCut(t string, f int) {
	data := paperCut{
		Type: t,
		Feed: f,
	}
	a := &action{
		Name: "paper_cut",
		Data: data,
	}
	g.actions = append(g.actions, a)
}
