package model

import (
	"fmt"
	"errors"
)

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
	data := `[
		{ “set_text_size”:3},
    	{ “printline” : “ร้านกาแฟ MOMO”},
    	{  “set_text_size”:1},
    	{ “printline” : “Ticketid  : 12”},
    	{ “printline” : “รายการสินค้า” },
    	{ “printline” : “ID     NAME      QTY     AMT”},
		{ “printline” : “2     Late        1       40.00”},
		{ “printline” : “----------------------”},
		{ “printline” : “รวมมูลค่าสินค้า     75”},
		{ “printline” : “เงินสด                100”},
		{ “printline” : “ขอบคุณที่ใช้บริการ”},
		{ “paper_cut”: {
          	"type": "full_cut",
            "feed": 1
      		}
        }
    ]
	`
	ch := make(chan *Message)
	m := &Message{
		Device: "printer",
		Command:"do_group",
		Data:   data,
	}
	H.Dev.Send <- m
	go func() {
		m2 := <-H.Dev.Send
		ch <- m2
	}()
	m = <-ch
	if !m.Result {
		return errors.New("Err: printer error.")
	}
	fmt.Println("Print success!")
	return nil
}

