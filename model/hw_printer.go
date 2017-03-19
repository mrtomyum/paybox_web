package model

import (
	"fmt"
	"errors"
	"github.com/gin-gonic/gin"
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

func (p *Printer) DoTicket(s *Sale) group {

	var g group
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
		Data:    data,
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

type group struct {
	actions []gin.H
}

func (g *group) print(t string) {
	a := gin.H{
		"action":      "print",
		"action_data": t,
	}
	g.actions = append(g.actions, a)
}

func (g *group) printLine(s string) {
	a := gin.H{
		"action":      "printline",
		"action_data": s,
	}
	g.actions = append(g.actions, a)
}

func (g *group) setTextSize(size int) {
	a := gin.H{
		"action":      "set_text_size",
		"action_data": size,
	}
	g.actions = append(g.actions, a)
}

func (g *group) newline() {
	a := gin.H{"action": "newline"}
	g.actions = append(g.actions, a)
}

func (g *group) printBarcode(t, d string) {
	data := gin.H{
		"type": t,
		"data": d,
	}
	a := gin.H{
		"action":      "print_barcode",
		"action_data": data,
	}
	g.actions = append(g.actions, a)
}

func (g *group) printQr(m, e int, t, d string) {
	data := gin.H{
		"mag":  m,
		"ect":  e,
		"type": t,
		"data": d,
	}
	a := gin.H{
		"action":      "print_qr",
		"action_data": data,
	}
	g.actions = append(g.actions, a)
}

func (g *group) paperCut(t string, f int) {
	data := gin.H{
		"type": t,
		"feed": f,
	}
	a := gin.H{
		"action":      "paper_cut",
		"action_data": data,
	}
	g.actions = append(g.actions, a)
}


