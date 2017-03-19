package model

import (
	"fmt"
	"errors"
	"strconv"
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
	case "no_paper": // Event แจ้งเตือนกระดาษหมดแล้ว
	}
}

func (p *Printer) DoTicket(s *Sale) doGroup {
	var g doGroup
	g.setTextSize(0)
	g.printLine("ร้านกาแฟสบายดี ยินดีต้อนรับ")
	sale := strconv.FormatFloat(s.Total, 'f', 2, 64)
	text := "มูลค่าขาย" + sale + "บาท"
	g.printLine(text)
	g.paperCut("full_cut", 80)
	fmt.Println(g.actions)
	return g
}

func (p *Printer) Print(s *Sale) error {
	fmt.Println("p.Print() run")
	data := p.DoTicket(s)
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

//===================================

type action struct {
	Name string `json:"action"`
	Data interface{} `json:"action_data"`
}

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
	a := &action{
		Name: "printline",
		Data: s,
	}
	g.actions = append(g.actions, a)
}

func (g *doGroup) setTextSize(size int) {
	a := &action{
		Name: "set_text_size",
		Data: size,
	}
	g.actions = append(g.actions, a)
}

func (g *doGroup) newline() {
	a := &action{Name: "newline"}
	g.actions = append(g.actions, a)
}

type barcode struct {
	Type string
	Data string
}

func (g *doGroup) printBarcode(t, d string) {
	data := &barcode{
		Type: t,
		Data: d,
	}
	a := &action{
		Name: "print_barcode",
		Data: data,
	}
	g.actions = append(g.actions, a)
}

type qrCode struct {
	Mag  int
	Ect  int
	Type string
	Data string
}

func (g *doGroup) printQr(m, e int, t, d string) {
	data := &qrCode{
		Mag:  m,
		Ect:  e,
		Type: t,
		Data: d,
	}
	a := &action{
		Name: "print_qr",
		Data: data,
	}
	g.actions = append(g.actions, a)
}

type paperCut struct {
	Type string
	Feed int
}

func (g *doGroup) paperCut(t string, f int) {
	data := &paperCut{
		Type: t,
		Feed: f,
	}
	a := &action{
		Name: "paper_cut",
		Data: data,
	}
	g.actions = append(g.actions, a)
}
