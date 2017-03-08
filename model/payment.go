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

func (pm *Payment) Pay(sale *Sale) error {
	// ตรวจสอบ Device Websocket Connecton ว่าสามารถติดต่อได้หรือยัง?
	pm.Remain = sale.Total
	if H.Dev == nil {
		log.Println("Device websocket ยังไม่ได้เชื่อมต่อเข้ามา")
	}
	// เปิดการรับชำระธนบัตร และ เหรียญ (Set Inhibit)
	fmt.Println("func Pay() -- 1. Start inhibit device BA, CA ")
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

		if PM.BillEscrow != 0 { // ถ้ารับธนบัตรให้ตรวจเงินทอน เพื่อคุมธนบัตรที่งดรับ
			fmt.Printf("3. รับธนบัตร Received Escrow = %v, Payment = %v Sale= %v\n", pm.BillEscrow, pm.Total, sale.Total)
			err := pm.RejectUnacceptedBill()
			if err != nil {
				return err
			}
		}

		if PM.Total >= sale.Total { // เมื่อชำระเงินครบหรือเกินกว่ายอดขายหรือไม่?
			fmt.Println("5. ยอดรับเงิน >= ยอดขาย")
			change := PM.Total - sale.Total
			if change != 0 { // หากต้องทอนเงิน (ไม่ต้องทอนให้ข้ามไป)
				fmt.Println("YES -> 6. ต้องทอนเงิน ")
				// ระบบจะยังไม่ Take เงิน ต้องตรวจก่อนว่ามีเหรียญพอทอนหรือไม่?
				if CB.Hopper >= change { // หากเหรียญใน Hopper พอทอน และยอดทอน != 0
					fmt.Println("YES -> 7. มีเหรียญพอทอน")
					if PM.BillEscrow != 0 { // ถ้ามีธนบัตรพักอยู่ ให้สั่งเก็บธนบัตร
						fmt.Println("YES -> 8.1 สั่งเก็บธนบัตร")
						err := BA.Take(true) // เก็บธนบัตรลงถัง
						if err != nil {
							return err
						}
					}
					fmt.Println("8.2 และทอนเหรียญ")
					err := CH.PayoutByCash(change) // Todo: เพิ่มกลไกวิเคราะห์เงินทอน แล้วสั่งทอนเป็นเหรียญ เพื่อป้องกันเหรียญหมด
					if err != nil {
						return err
						log.Println("Error on CH Payout():", err.Error())
					}
					fmt.Println("SUCCESS -- ทอนเหรียญจาก Hopper สำเร็จ PM.Total=", PM.Total)
				} else {
					fmt.Println("NO -> 9 #เหรียญไม่พอทอน# รับธนบัตรรึเปล่า")
					if PM.BillEscrow != 0 { // ถ้ามียอดรับล่าสุดเป็นธนบัตร (ที่ถูกพักไว้)
						fmt.Println("YES -> 9.1 ถ้ารับด้วยธนบัตรให้คายธนบัตรคืนลูกค้า -- สั่งคายธนบัตร")
						err := BA.Take(false) // คายธนบัตร (Reject)
						if err != nil {
							return err
						}
						fmt.Println("SUCCESS -- คายธนบัตรเมื่อเหรียญใน Hopper ไม่พอทอน PM.Total=", PM.Total)
					}
					fmt.Println("No -> 9.2 รับมาด้วยเหรียญ -- ให้ทอนเหรียญตามจำนวนที่รับมา")
					err := CH.PayoutByCash(PM.CoinEscrow)
					if err != nil {
						return err
						log.Println("Error on CH Payout():", err.Error())
					}
				}
			} else {
				fmt.Println("NO -> 10 รับด้วยธนบัตรรึเปล่า?")
				if PM.BillEscrow != 0 { // เฉพาะธนบัตรต้องสั่ง Take ก่อน
					// กินธนบัตรที่พักไว้ *ระวัง! ถ้า Dev client ยังไม่เปิดคอนเนคชั่นจะ runtime error: invalid memory address or nil pointer derefere
					err := BA.Take(true) // เก็บธนบัตรลงถัง
					if err != nil {
						return err
					}
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
	if PM.BillEscrow != 0 { // ถ้ามีธนบัติพักไว้
		err := BA.Take(true) // เก็บธนบัตรลงถัง
		if err != nil {
			return err
		}
		BA.Stop()
	}
	// Success
	PM.Coin = PM.Total - PM.Bill
	PM.Total = PM.Coin
	PM.Bill = 0

	// CoinHopper สั่งให้จ่ายเหรียญที่คงค้างตามยอดคงเหลือ PM.Coin ออกด้านหน้า
	err := CH.PayoutByCash(PM.Coin)
	if err != nil {
		return err
	}
	PM.Total = 0 // เคลียร์ยอดเงินค้างให้หมด
	PM.BillEscrow = 0
	PM.CoinEscrow = 0

	// Send message response back to Web Client
	c.Msg.Type = "response"
	c.Msg.Result = true
	c.Msg.Data = "sucess"
	c.Send <- c.Msg
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
		Device: "host",
		Command:"accepted_bill",
		Type:   "event",
		Data:   AB,
	}
	fmt.Println("Send message to WebUI = ", m)
	H.Web.Send <- m
}

func (pm *Payment) RejectUnacceptedBill() error {
	fmt.Println("4. ถ้ารับธนบัตร ตรวจสอบเพื่อ Reject ธนบัตรที่ไม่รับ")
	switch PM.BillEscrow {
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