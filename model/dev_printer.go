package model

import "fmt"

type Printer struct {
	Id     string
	Status string
	Send   chan *Message
}

func (p *Printer) Event(c *Client) {
	switch c.Msg.Command {
	case "machine_id": // ร้องขอหมายเลข Serial Number ของ อุปกรณ์ Printer
	case "do_single":  //ส่ังการเคร่ืองปริ้นเตอร์ แบบส่งคาส่ังการกระทาคาสั่งเดียว โดย action_name และ action_data สามารถดูได้จากตาราง Action
	case "do_group":   //ส่ังการเคร่ืองปร้ินเตอร์ แบบส่งคาส่ังการกระทาแบบเปน็ ชุด โดย action_name และ action_data สามารถดูได้จากตาราง Action
	case "near_end":   // Event แจ้งเตือนกระดาษใกล้หมด
	case "no_paper":   // Event แจ้งเตือนกระดาษหมดแล้ว
	}
}

func (p *Printer) Print(s *Sale) error {
	fmt.Println("p.Print() run")
	data :=
	return nil
}

