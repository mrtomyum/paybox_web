package model

import (
	"fmt"
	"log"
	"errors"
)

// Payment คือยอดเงินพัก ยังไม่ได้รับชำระ
type Payment struct {
	Coin   float64 // มูลค่าเหรียญที่รับมา ที่ยังไม่ได้รับชำระ
	Bill   float64 // มูลค่าธนบัตรที่รับมา ที่ยังไม่ได้รับชำระ
	Escrow float64 // มูลค่าธนบัตรที่พักอยู่ในตัวรับธนบัตร BA
	//Card  float64 // มูลค่าบัตรเครดิตที่รับชำระแล้ว
	Total float64 // มูลค่าเงินพักทั้งหมด
	Send  chan *Message
}

// CashBox คือถังเก็บเงิน แยกเป็น 3 จุด
// คือ Hopper ถังเก็บเหรียญ CoinBox และถังเก็บธนบัตร BillBox
type CashBox struct {
	Hopper float64 // มูลค่าเหรียญใน Coin Hopper
	Coin   float64 // มูลค่าเหรียญใน CainBox
	Bill   float64 // มูลค่าธนบัตรในกล่องเก็บธนบัตร
	Total  float64 // รวมมูลค่าเงินในตู้นี้
}

// AcceptedValue ระบุค่ายอดขายขั้นต่ำที่ยอมรับธนบัตรแต่ละขนาด 0 = ไม่จำกัด
type AcceptedValue struct {
	B20   int `json:"b20"`
	B50   int `json:"b50"`
	B100  int `json:"b100"`
	B500  int `json:"b500"`
	B1000 int `json:"b1000"`
}

type AcceptedBill struct {
	B20   bool `json:"b20"`
	B50   bool `json:"b50"`
	B100  bool `json:"b100"`
	B500  bool `json:"b500"`
	B1000 bool `json:"b1000"`
}

func (pm *Payment) Pay(sale *Sale) error {
	// เปิดการรับชำระธนบัตร และ เหรียญ (Set Inhibit)
	fmt.Printf("func Pay() -- \n1. Start inhibit device BA, CA \n")
	CA.Start()
	BA.Start()

	// หากธนบัตร หรือเหรียญที่ชำระยังมีมูลค่าน้อยกว่ายอดขาย (Payment < Sale)
	// ระบบจะ Take เงิน และจะสะสมยอดรับชำระ และส่ง command: "onhand" เป็น event กลับตลอดเวลาจนกว่าจะได้ยอด Payment >= Sale
	for {
		fmt.Println("2. Waiting payment form BA or CA")
		m := <-PM.Send
		fmt.Printf("3. Received Escrow = %v, Payment = %v Sale= %v\n", PM.Escrow, PM.Total, S.Total)

		if PM.Total >= sale.Total { // เมื่อชำระเงินครบหรือเกิน
			break
		}
		if m.Device == "bill_acc" { // เฉพาะธนบัตรต้องสั่ง Take ก่อน
			// กินธนบัตรที่พักไว้ *ระวัง! ถ้า Dev client ยังไม่เปิดคอนเนคชั่นจะ runtime error: invalid memory address or nil pointer derefere
			err := BA.Take(true)
			if err != nil {
				return err
			}
		}
	}

	change := PM.Total - sale.Total
	// ระบบจะยังไม่ Take เงิน ต้องตรวจก่อนว่ามีเหรียญพอทอนหรือไม่?
	if CB.Hopper < change { // หากเหรียญใน Hopper ไม่พอทอน
		if PM.Escrow != 0 { // มียอดรับล่าสุดเป็นธนบัตร (ที่ถูกพักไว้)
			err := BA.Take(false) // คายธนบัตร (Reject)
			if err != nil {
				return err
			}
			fmt.Println("คายธนบัตรเมื่อเหรียญใน Hopper ไม่พอทอน PM.Total=", PM.Total)
		}
	} else {
		// เงินไม่พอทอน ต้อง Reject และ เข้า Loop Payment ใหม่
	}

	// ทอนเงินจาก CoinHopper ถ้ามี
	if change > 0 {
		err := CH.PayoutByCash(change) // Todo: เพิ่มกลไกวิเคราะห์เงินทอน แล้วสั่งทอนเป็นเหรียญ เพื่อป้องกันเหรียญหมด
		if err != nil {
			return err
			log.Println("Error on CH Payout():", err.Error())
		}
	}

	// ปิดการรับชำระที่อุปกรณ์
	CA.Stop()
	BA.Stop()
	return nil
}

// OnHand ส่งค่าเงินพัก Escrow ไว้กลับไปให้ web
func (pm *Payment) OnHand(web *Client) {
	fmt.Println("method *Host.OnHand()...")
	web.Msg.Command = "onhand"
	web.Msg.Result = true
	web.Msg.Type = "event"
	web.Msg.Data = pm.Total
	web.Send <- web.Msg
}

// Cancel คืนเงินจากทุก Device โดยตรวจสอบเงิน Escrow ใน Bill Acceptor ด้วยถ้ามีให้คืนเงิน
func (pm *Payment) Cancel(c *Client) error {
	fmt.Println("Host.Cancel()...")

	// Check Bill Acceptor
	if PM.Total == 0 { // ไม่มีเงินพัก
		log.Println("ไม่มีเงินพัก:")
		c.Msg.Type = "response"
		c.Msg.Result = false
		c.Msg.Data = "ไม่มีเงินพัก"
		c.Send <- c.Msg
		return errors.New("ไม่มีเงินพัก")
	}
	// สั่งให้ BillAcceptor คืนเงินที่พักไว้
	m1 := &Message{
		Device:  "bill_acc",
		Command: "escrow",
		Type:    "request",
		Result:  true,
		Data:    false,
	}
	H.Dev.Send <- m1

	// Check BillAcc response
	err := H.Dev.Ws.ReadJSON(&m1)
	if err != nil {
		log.Println("Host.Cancel() error ->", m1.Data)
		return err
	}

	// Success
	PM.Coin = PM.Total - PM.Bill
	PM.Bill = 0

	// CoinHopper สั่งให้จ่ายเหรียญที่คงค้างตามยอด coinHopperEscrow ออกด้านหน้า
	m2 := &Message{
		Device:  "coin_hopper",
		Command: "payout_by_cash",
		Type:    "request",
		Data:    PM.Coin,
	}
	H.Dev.Send <- m2

	// Check if error from CoinHopper
	err = H.Dev.Ws.ReadJSON(&m2)
	if err != nil {
		log.Println("Cancel() Coin Hopper error:", err)
		c.Msg.Result = false
		c.Msg.Type = "response"
		c.Msg.Data = m2.Data
		c.Send <- c.Msg
		return err
	}
	PM.Total = 0 // เคลียร์ยอดเงินค้างให้หมด

	// Send message response back to Web Client
	c.Msg.Type = "response"
	c.Msg.Result = true
	c.Msg.Data = "sucess"
	c.Send <- c.Msg
	return nil
}

func CheckAcceptedBill(s *Sale) {
	switch {
	case s.Total <= AV.B20:
		AB.B20 = false
		fallthrough
	case s.Total <= AV.B50:
		AB.B50 = false
		fallthrough
	case s.Total >= AV.B100:
		AB.B100 = false
		fallthrough
	case s.Total >= AV.B500:
		AB.B500 = false
		fallthrough
	case s.Total >= AV.B1000:
		AB.B1000 = false
		fallthrough
	}
}