package model

import (
	"fmt"
	"log"
	"errors"
)

// Payment คือยอดเงินพัก ยังไม่ได้รับชำระ
type Payment struct {
	Coin       float64 // มูลค่าเหรียญที่รับมาทั้งหมด แต่ยังไม่ได้รับชำระ
	CoinEscrow float64 // มูลค่าเหรียญที่พัก
	Bill       float64 // มูลค่าธนบัตรที่รับมาทั้งหมด แต่ยังไม่ได้รับชำระ
	BillEscrow float64 // มูลค่าธนบัตรที่พักอยู่ในตัวรับธนบัตร BA
	Total      float64 // มูลค่าเงินพักทั้งหมด
	Remain     float64 // เงินคงค้างชำระ
	Received   chan *Message
	//Card  float64 // มูลค่าบัตรเครดิตที่รับชำระแล้ว
}

// CashBox คือถังเก็บเงิน แยกเป็น 3 จุด
// คือ ชุดทอนเหรียญ Hopper, ถังเก็บเหรียญล้น CoinBox และถังเก็บธนบัตร BillBox
type CashBox struct {
	Hopper float64 // มูลค่าเหรียญใน Coin Hopper
	Coin   float64 // มูลค่าเหรียญใน CainBox
	Bill   float64 // มูลค่าธนบัตรในกล่องเก็บธนบัตร
	Total  float64 // รวมมูลค่าเงินในตู้นี้
}

// AcceptedValue ระบุค่ายอดขายขั้นต่ำที่ยอมรับธนบัตรแต่ละขนาด 0 = ไม่จำกัด
type AcceptedValue struct {
	B20   float64 `json:"b20"`
	B50   float64 `json:"b50"`
	B100  float64 `json:"b100"`
	B500  float64 `json:"b500"`
	B1000 float64 `json:"b1000"`
}

type AcceptedBill struct {
	B20   bool `json:"b20"`
	B50   bool `json:"b50"`
	B100  bool `json:"b100"`
	B500  bool `json:"b500"`
	B1000 bool `json:"b1000"`
}

// *Payment Pay() ทำหน้าที่จัดการกระบวนการรับเงิน ทอนเงิน ให้สมบูรณ์
func (pm *Payment) Pay(sale *Sale) error {
	// ตรวจสอบ WebSocket Connection?
	if H.Dev == nil || H.Web == nil {
		log.Println("HW_SERVICE หรือ WebUI websocket ยังไม่ได้เชื่อมต่อ")
	}

	pm.Remain = sale.Total
	// เปิดการรับชำระธนบัตร และ เหรียญ (Set Inhibit)
	fmt.Println("func Pay() -- 1. Start Payment device:")
	CA.Start()
	BA.Start()

	// หากธนบัตร หรือเหรียญที่ชำระยังมีมูลค่าน้อยกว่ายอดขาย (Payment < Sale)
	// ระบบจะ Take เงิน และจะสะสมยอดรับชำระ และส่ง command: "onhand" เป็น event กลับตลอดเวลา
	// จนกว่าจะได้ยอด Payment >= Sale จึงค่อย break ออกจาก for loop
	for {
		pm.CheckAcceptedBill(sale)
		pm.DisplayAcceptedBill() // DisplayAcceptedBill() ส่งรายการธนบัตรที่รับได้ไปแสดงบนหน้าจอ
		fmt.Println("2. Waiting payment form BA or CA")
		<-pm.Received // Waiting for message from payment device.

		fmt.Printf("3. รับธนบัตร Received Escrow = %v, Payment = %v Sale= %v\n", pm.BillEscrow, pm.Total, sale.Total)
		switch {
		case PM.Total >= sale.Total:
			switch {
			case pm.BillEscrow != 0:
				err := pm.RejectUnacceptedBill()
				if err != nil {
					return err
				}
			case pm.CoinEscrow != 0:

			}
		}

		if PM.Total >= sale.Total { // เมื่อชำระเงินครบหรือเกินกว่ายอดขายหรือไม่?
			fmt.Println("5. ยอดรับเงิน >= ยอดขาย")
			change := PM.Total - sale.Total

			if change != 0 { // หากต้องทอนเงิน (ไม่ต้องทอนให้ข้ามไป)
				fmt.Println("YES -> 6. ต้องทอนเงิน ")
				// ระบบจะยังไม่ Take เงิน ต้องตรวจก่อนว่ามีเหรียญพอทอนหรือไม่?
				if CB.Hopper < change { // หากเหรียญใน Hopper ไม่พอทอน และยอดทอน != 0
					err := pm.coinShortage()
					if err != nil {
						return err
					}
				}
				fmt.Println("YES -> 7. มีเหรียญพอทอน")
				err := BA.Take(true) // เก็บธนบัตรลงถัง
				if err != nil {
					return err
				}
				fmt.Println("YES -> 8.1 สั่งเก็บธนบัตรสำเร็จ")

				err = CH.PayoutByCash(change) // Todo: เพิ่มกลไกวิเคราะห์เงินทอน แล้วสั่งทอนเป็นเหรียญ เพื่อป้องกันเหรียญหมด
				if err != nil {
					return err
					log.Println("Error on CH Payout():", err.Error())
				}
				fmt.Println("8.2 SUCCESS -- ทอนเหรียญจาก Hopper สำเร็จ PM.Total=", PM.Total)

			} else { // ไม่ต้องทอนเงิน
				if pm.BillEscrow != 0 { // ถ้ารับธนบัตร
					err := BA.Take(true) // ให้เก็บธนบัตรลงถัง
					if err != nil {
						return err
					}
					fmt.Println("เก็บธนบัตรสำเร็จ")
				}
			}
		}
		// เช็คอีกรอบ ถ้ายอดเงินรับ >= ขอดขาย แล้ว ให้ล้างข้อมูล
		// และ break ออกจาก for(ever) loop
		if PM.Total >= sale.Total {
			PM.Total = 0
			PM.Coin = 0
			PM.Bill = 0
			PM.Remain = 0
			PM.BillEscrow = 0
			PM.CoinEscrow = 0
			break
		}
	}

	// ปิดการรับชำระที่อุปกรณ์
	CA.Stop()
	BA.Stop()

	m := &Message{
		Device:  "host",
		Command: "payment",
		Type:    "event",
		Data:    "success",
	}

	H.Web.Send <- m
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

	// ตรวจสอบก่อนว่าหากคืนธนบัตรใบล่าสุดใบเดียว เหรียญใน Hopper จะพอคืนตามยอดเงินรับชำระหรือไม่?
	if pm.Total-pm.BillEscrow > CB.Hopper {
		// ให้แจ้งเตือนการ Cancel ล้มเหลว
		// แล้วคืนธนบัตรใบล่าสุด พร้อมทั้งพิมพ์คูปองคืนเงิน
		// โดยไม่ทอนเหรียญที่เหลือเนื่องจากหากมีเหรียญเหลือน้อยมักมีปัญหาในการทอนล่าช้า
		// หรืออาจมีจำนวนเหรียญไม่ตรงกับที่ได้รับแจ้งตอนเริ่มระบบ
		pm.refund(PM.Total, PM.BillEscrow)
	}

	// Check Bill Acceptor
	if PM.Total == 0 { // ไม่มีเงินพัก
		log.Println("ไม่มีเงินพัก:")
		c.Msg.Type = "response"
		c.Msg.Result = false
		c.Msg.Data = "ไม่มีเงินพัก"
		c.Send <- c.Msg
		return errors.New("ไม่มีเงินพัก")
	}

	// สั่งให้ BillAcceptor คืนเงินที่พักไว้ ซึ่งจะคืนได้เพียงใบล่าสุด
	err := BA.Take(true) // เก็บธนบัตรลงถัง
	if err != nil {
		return err
	}
	BA.Stop()
	// Success
	PM.Coin = PM.Total - PM.Bill
	PM.Total = PM.Coin
	PM.Bill = 0

	// CoinHopper สั่งให้จ่ายเหรียญที่คงค้างตามยอดคงเหลือ PM.Coin ออกด้านหน้า
	err = CH.PayoutByCash(PM.Coin)
	if err != nil {
		return err
	}

	// Send message response back to Web Client
	c.Msg.Type = "response"
	c.Msg.Result = true
	c.Msg.Data = PM.Coin
	c.Send <- c.Msg

	pm.Total = 0 // เคลียร์ยอดเงินค้างให้หมด
	pm.Bill = 0
	pm.BillEscrow = 0
	pm.Coin = 0
	pm.CoinEscrow = 0
	pm.Remain = 0
	return nil
}

func (pm *Payment) CheckAcceptedBill(s *Sale) {
	// ตรวจยอดขาย เทียบกับ เงินทอนใน Hopper ว่าพอหรือไม่
	// เพื่อเลือกเปิด/ปิดรับธนบัตร เงื่อนไขคือ...
	// ธนบัตรที่จะรับ ถ้าหักยอดคงค้างชำระ แล้วต้องน้อยกว่า เงินที่เหลือใน Hopper
	switch {
	case s.Total < AV.B20 || 20-pm.Remain > CB.Hopper:
		AB.B20 = false
		fallthrough
	case s.Total < AV.B50 || 50-pm.Remain > CB.Hopper:
		AB.B50 = false
		fallthrough
	case s.Total < AV.B100 || 100-pm.Remain > CB.Hopper:
		AB.B100 = false
		fallthrough
	case s.Total < AV.B500 || 500-pm.Remain > CB.Hopper:
		AB.B500 = false
		fallthrough
	case s.Total < AV.B1000 || 1000-pm.Remain > CB.Hopper:
		AB.B1000 = false
	}
}

func (pm *Payment) DisplayAcceptedBill() {
	// Check MinAcceptedBill500 & 1000
	m := &Message{
		Device:  "host",
		Command: "accepted_bill",
		Type:    "event",
		Data:    AB,
	}
	fmt.Println("Send message to WebUI = ", m)
	H.Web.Send <- m
}

func (pm *Payment) RejectUnacceptedBill() error {
	if pm.BillEscrow != 0 { // ถ้ารับธนบัตรให้ตรวจเงินทอน เพื่อคุมธนบัตรที่งดรับ
		return ErrNoBillEscrow
	}
	fmt.Println("4. ถ้ารับธนบัตร ตรวจสอบเพื่อ Reject ธนบัตรที่ไม่รับ")
	switch pm.BillEscrow {
	case 20.0:
		if !AB.B20 {
			BA.Take(false)
		}
	case 50.0:
		if !AB.B50 {
			BA.Take(false)
		}
	case 100.0:
		if !AB.B100 {
			BA.Take(false)
		}
	case 500:
		if !AB.B500 {
			BA.Take(false)
		}
	case 1000:
		if !AB.B1000 {
			BA.Take(false)
		}
	default:
		fmt.Println("ไม่เข้าเงื่อนไข")
	}
	fmt.Println("PM.BillEscrow =", pm.BillEscrow)
	fmt.Println("AcceptedBill = ", AB)
	return nil
}

func (pm *Payment) coinShortage() error {
	fmt.Println("NO -> 9 รับธนบัตรรึเปล่า")
	if pm.BillEscrow != 0 { // ถ้ามียอดรับล่าสุดเป็นธนบัตร (ที่ถูกพักไว้)
		fmt.Println("YES -> 9.1 ถ้ารับด้วยธนบัตรให้คายธนบัตรคืนลูกค้า -- สั่งคายธนบัตร")
		err := BA.Take(false) // คายธนบัตร (Reject)
		if err != nil {
			return err
		}
		fmt.Println("SUCCESS -- คายธนบัตรเมื่อเหรียญใน Hopper ไม่พอทอน PM.Total=", PM.Total)
	}
	fmt.Println("No -> 9.2 รับมาด้วยเหรียญ -- ให้คืนเหรียญตามจำนวนที่รับมา")
	err := CH.PayoutByCash(pm.CoinEscrow)
	if err != nil {
		return err
		log.Println("Error on CH Payout():", err.Error())
	}
	return nil
}

func (pm *Payment) refund(total, billEscrow float64) error {
	err := BA.Take(false)
	if err != nil {
		return err
	}
	// Print ใบคืนเงิน (Refund) ตามยอดเงินคงเหลือ
	rf := total - billEscrow
	err = P.makeRefund(rf)
	if err != nil {
		return err
	}
	return nil
}
