package model

type CoinAcceptor struct {
	Id     string
	Status string
	Send   chan *Message
}

func (ca *CoinAcceptor) Event(c *Client) {
	switch c.Msg.Command {
	case "machine_id":        // ร้องขอหมายเลข Serial Number ของ อุปกรณ์ Coins Acceptor
	case "inhibit":           // ร้องขอ สถานะ Inhibit (รับ-ไม่รับเหรียญ) ของ Coins Acceptor
	case "set_inhibit":       // ตั้งค่า Inhibit (รับ-ไม่รับเหรียญ) ของ Coins Acceptor
	case "recently_inserted": // ร้องขอจานวนเงินของเหรียญล่าสุดที่ได้รับ
	case "received":          // Event น้ีจะเกิดขึ้นเมื่อเคร่ืองรับเหรียญได้รับเหรียญ
	}
}