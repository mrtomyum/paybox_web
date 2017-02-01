package model

import (
	"fmt"
	"errors"
)

type Printer struct {
	machineId string `json:"machine_id"`
	Status    string
	Send      chan *Message
}

func (p *Printer) Event(c *Client) {
	switch c.Msg.Command {
	case "machine_id", "do_single", "do_group":
		p.Send <- c.Msg
	//case "machine_id": // ร้องขอหมายเลข Serial Number ของ อุปกรณ์ Printer
	//case "do_single":  //ส่ังการเคร่ืองปริ้นเตอร์ แบบส่งคาส่ังการกระทาคาสั่งเดียว โดย action_name และ action_data สามารถดูได้จากตาราง Action
	//case "do_group":   //ส่ังการเคร่ืองปร้ินเตอร์ แบบส่งคาส่ังการกระทาแบบเปน็ ชุด โดย action_name และ action_data สามารถดูได้จากตาราง Action
	case "near_end":   // Event แจ้งเตือนกระดาษใกล้หมด
	case "no_paper":   // Event แจ้งเตือนกระดาษหมดแล้ว
	}
}

func (p *Printer) Print(s *Sale) error {
	fmt.Println("p.Print() run")
	data, err := p.makeSaleSlip(s)
	if err != nil {
		return err
	}
	ch := make(chan *Message)
	m := &Message{
		Device: "printer",
		Command:"do_group",
		Data:   data,
	}
	H.Dev.Send <- m
	fmt.Println("1. สั่งพิมพ์ รอ Priner ตอบสนอง")
	go func() {
		m2 := <-p.Send
		ch <- m2
	}()
	m = <-ch
	if !m.Result {
		return errors.New("Err: printer error.")
	}
	fmt.Println("พิมพ์สำเร็จ Print success!")
	return nil
}

func (p *Printer) makeSaleSlip(s *Sale) (data string, err error) {
	header := `[
		{"set_text_size":3},
		{"printline" : "ร้านกาแฟ MOMO"},
		{"set_text_size":1},
		{"printline": "Ticketid  : 12"},
		{"printline": "ID     NAME      QTY     AMT"},
	`
	item := `
		{"printline": "2     Late        1       40.00"},

	`
	footer := `
		{"printline": "----------------------"},
		{"printline": "รวมมูลค่าสินค้า     %v"},
		{"printline": "รับเงิน           %v"},
		{"printline": "เงินทอน          %v"},
		{"printline": "ขอบคุณที่ใช้บริการ"},
		{"paper_cut": {"type": "partial_cut","feed": 1}}
	`
	queue := `
		{"printline": "Ticket" },
		{"set_text_size":8},
		{%s},
		{"paper_cut": {"type": "full_cut","feed": 1}}
		]
	`
	data = fmt.Sprintln(header+item+footer+queue, s.Total, s.Pay, s.Change)
	fmt.Println("data=", data)
	return data, nil
}

func (p *Printer) makeTicket(s *Sale) {

}