package model

import (
	"fmt"
	"log"
	"errors"
)

var ErrCoinShortage error = errors.New("Not enough coin to change// ไม่มีเหรียญพอทอน")

// Payment คือยอดเงินพัก ยังไม่ได้รับชำระ
type Payment struct {
	coin       float64 // มูลค่าเหรียญที่รับมาทั้งหมด แต่ยังไม่ได้รับชำระ
	bill       float64 // มูลค่าธนบัตรที่รับมาทั้งหมด แต่ยังไม่ได้รับชำระ
	billEscrow float64 // มูลค่าธนบัตรที่พักอยู่ในตัวรับธนบัตร BA
	total      float64 // มูลค่าเงินพักทั้งหมด
	remain     float64 // เงินคงค้างชำระ
	receivedCh chan *Message
	//Card  float64 // มูลค่าบัตรเครดิตที่รับชำระแล้ว
}

// CashBox คือถังเก็บเงิน แยกเป็น 3 จุด
// คือ ชุดทอนเหรียญ hopper, ถังเก็บเหรียญล้น CoinBox และถังเก็บธนบัตร BillBox
type CashBox struct {
	hopper float64 // มูลค่าเหรียญใน coin hopper
	coin   float64 // มูลค่าเหรียญใน CainBox
	bill   float64 // มูลค่าธนบัตรในกล่องเก็บธนบัตร
	total  float64 // รวมมูลค่าเงินในตู้นี้
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

// init() ทำการรีเซ็ทค่าที่ควรถูกตั้งใหม่ทุกครั้งที่สร้าง Payment ใหม่ขึ้นมา
func (pm *Payment) init() {
	AB = &AcceptedBill{
		B20:   true,
		B50:   true,
		B100:  true,
		B500:  true,
		B1000: true,
	}
	// เปิดการรับชำระธนบัตร และ เหรียญ (Set Inhibit)
	fmt.Println("func New() -- 1. Start Payment device:")
	CA.Start()
	BA.Start()
}

// *Payment New() ทำหน้าที่จัดการกระบวนการรับเงิน ทอนเงิน ให้สมบูรณ์
func (pm *Payment) New(sale *Sale) error {
	// ตรวจสอบ WebSocket Connection?
	if H.Hw == nil || H.Web == nil {
		log.Println("HW_SERVICE หรือ WebUI websocket ยังไม่ได้เชื่อมต่อ")
	}
	if sale.Total == 0 {
		return errors.New("Sale Total is 0 cannot do payment.")
	}
	pm.init()
	pm.remain = sale.Total

	// หากธนบัตร หรือเหรียญที่ชำระยังมีมูลค่าน้อยกว่ายอดขาย (Payment < Sale)
	// ระบบจะ Take เงิน และจะสะสมยอดรับชำระ และส่ง command: "onhand" เป็น event กลับตลอดเวลา
	// จนกว่าจะได้ยอด Payment >= Sale
	for pm.total < sale.Total {
		pm.checkAcceptedBill(sale)
		pm.displayAcceptedBill() // displayAcceptedBill() ส่งรายการธนบัตรที่รับได้ไปแสดงบนหน้าจอ
		fmt.Println("1. Waiting payment form BA or CA")
		msg := <-pm.receivedCh // Waiting for message from payment device.

		fmt.Printf("2. ยอดขาย = %v รับจาก = %v, Payment = %v \n", sale.Total, msg.Device, msg.Data)
		if msg.Device == "bill_acc" { //ถ้าเป็นธนบัตร
			err := pm.rejectUnacceptedBill()
			if err != nil {
				return err
			}
			err = BA.Take(true) // ให้เก็บธนบัตรลงถัง
			if err != nil {
				return err
			}
			fmt.Printf("เก็บธนบัตรสำเร็จ: pm.total= %v sale.total= %v pm.remain= %v", pm.total, sale.Total, pm.remain)
			pm.sendOnHand(H.Web)
		}
	}
	change := pm.total - sale.Total
	if change > 0 { //ถ้าต้องทอนเงิน
		fmt.Println("pm.change()")
		err := pm.change(change)
		if err != nil {
			return err
		}
	}

	// ล้างข้อมูล
	fmt.Println("Reset Payment: ล้างยอดรับชำระ")
	PM.coin = 0
	PM.bill = 0
	PM.billEscrow = 0
	PM.remain = 0
	PM.total = 0
	// ส่งยอดที่ล้างแล้วให้ WebUI
	//pm.sendOnHand(H.Web)
	// ปิดการรับชำระที่อุปกรณ์
	CA.Stop()
	BA.Stop()

	m := &Message{
		Device:  "web",
		Command: "payment",
		Type:    "event",
		Data:    "success",
	}
	H.Web.Send <- m
	return nil
}

// OnHand ส่งค่าเงินพัก Escrow ไว้กลับไปให้ web
func (pm *Payment) sendOnHand(web *Socket) {
	fmt.Println("method *Host.sendOnHand()...")
	web.Msg.Command = "onhand"
	web.Msg.Result = true
	web.Msg.Type = "event"
	web.Msg.Data = pm.total
	web.Send <- web.Msg
}

// Cancel คืนเงินจากทุก Device โดยตรวจสอบเงิน Escrow ใน bill Acceptor ด้วยถ้ามีให้คืนเงิน
func (pm *Payment) Cancel(c *Socket) {
	fmt.Println("call *Payment.Cancel()")

	// ตรวจสอบก่อนว่าหากคืนธนบัตรใบล่าสุดใบเดียว เหรียญใน hopper จะพอคืนตามยอดเงินรับชำระหรือไม่?
	//if pm.total-pm.billEscrow > CB.hopper {
	//	// คืนธนบัตรใบล่าสุด พร้อมทั้งพิมพ์คูปองคืนเงิน
	//	// โดยไม่ทอนเหรียญที่เหลือเนื่องจากหากมีเหรียญเหลือน้อยมักมีปัญหาในการทอนล่าช้า
	//	// หรืออาจมีจำนวนเหรียญไม่ตรงกับที่ได้รับแจ้งตอนเริ่มระบบ
	//	// ตอนนี้เลือกวิธีให้ยกเลิกการขาย
	//	pm.refund(PM.total, PM.billEscrow)
	//	return ErrCoinShortage
	//}
	fmt.Printf("pm.billEscrow: %v pm.total: %v ", pm.billEscrow, pm.total)
	// Check bill Acceptor

	if pm.total == 0 { // ไม่มีเงินรับชำระ
		if pm.billEscrow != 0 { //มีธนบัตร
			// สั่งให้ BillAcceptor คืนเงินที่พักไว้ ซึ่งจะคืนได้เพียงใบล่าสุด
			err := BA.Take(false) // คายธนบัตร
			if err != nil {
				log.Println(err.Error())
			}
		}
		BA.Stop()
		CA.Stop()
		c.Msg.Type = "response"
		c.Msg.Result = false
		c.Msg.Data = "ไม่มีเงินรับ"
		c.Send <- c.Msg
		pm.reset()
		return
	}
	BA.Stop()
	CA.Stop()

	// CoinHopper สั่งให้จ่ายเหรียญที่คงค้างตามยอดคงเหลือ PM.coin ออกด้านหน้า
	change := PM.total - PM.billEscrow
	err := CH.PayoutByCash(change)
	if err != nil {
		log.Println(err.Error())
	}

	// Send message response back to Web Socket
	c.Msg.Type = "response"
	c.Msg.Result = true
	c.Msg.Data = PM.coin
	c.Send <- c.Msg
	pm.reset()
}

func (pm *Payment) checkAcceptedBill(s *Sale) {
	// ตรวจยอดขาย เทียบกับ เงินทอนใน hopper ว่าพอหรือไม่
	// เพื่อเลือกเปิด/ปิดรับธนบัตร เงื่อนไขคือ...
	// ธนบัตรที่จะรับ ถ้าหักยอดคงค้างชำระ แล้วต้องน้อยกว่า เงินที่เหลือใน hopper
	switch {
	case s.Total < AV.B20 || 20-pm.remain > CB.hopper:
		AB.B20 = false
		fallthrough
	case s.Total < AV.B50 || 50-pm.remain > CB.hopper:
		AB.B50 = false
		fallthrough
	case s.Total < AV.B100 || 100-pm.remain > CB.hopper:
		AB.B100 = false
		fallthrough
	case s.Total < AV.B500 || 500-pm.remain > CB.hopper:
		AB.B500 = false
		fallthrough
	case s.Total < AV.B1000 || 1000-pm.remain > CB.hopper:
		AB.B1000 = false
	}
}

func (pm *Payment) displayAcceptedBill() {
	// Check MinAcceptedBill500 & 1000
	m := &Message{
		Device:  "web",
		Command: "accepted_bill",
		Type:    "event",
		Data:    AB,
	}
	fmt.Println("Send message to WebUI = ", m)
	H.Web.Send <- m
}

// ตรวจสอบธนบัตรที่ต้อง  Reject
func (pm *Payment) rejectUnacceptedBill() error {
	fmt.Println("4. ถ้ารับธนบัตร ตรวจสอบเพื่อ Reject ธนบัตรที่ไม่รับ")
	if pm.billEscrow == 0 {
		log.Println(ErrCoinShortage.Error())
		return ErrNoBillEscrow
	}
	switch pm.billEscrow {
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
	fmt.Println("PM.billEscrow =", pm.billEscrow, "AcceptedBill = ", AB)
	return nil
}

// refund() เมื่อต้องการพิมพ์ใบคืนเงิน ในกรณีเงินทอนไม่พอ
func (pm *Payment) refund(total, billEscrow float64) error {
	err := BA.Take(false) //คายธนบัตร
	if err != nil {
		return err
	}

	// Todo: Print ใบคืนเงิน (Refund) ตามยอดเงินคงเหลือ
	rf := total - billEscrow
	err = P.makeRefund(rf) //ยังไม่เสร็จ
	if err != nil {
		return err
	}
	return nil
}

func (pm *Payment) change(value float64) error {
	fmt.Println("YES -> 6. ต้องทอนเงิน ")
	// ส่ง Web Socket: "command": "change", "data": value
	m := &Message{
		Device:  "web",
		Type:    "event",
		Command: "change",
		Data:    value,
	}
	H.Web.Send <- m
	// ระบบจะยังไม่ Take เงิน ต้องตรวจก่อนว่ามีเหรียญพอทอนหรือไม่?
	//if CB.hopper < value { // หากเหรียญใน hopper ไม่พอทอน และยอดทอน != 0
	//	err := pm.coinShortage()
	//	if err != nil {
	//		return err
	//	}
	//	return ErrCoinShortage
	//}
	fmt.Println("YES -> 7. มีเหรียญพอทอน")
	err := CH.PayoutByCash(value) // Todo: เพิ่มกลไกวิเคราะห์เงินทอน แล้วสั่งทอนเป็นเหรียญ เพื่อป้องกันเหรียญหมด
	if err != nil {
		return err
		log.Println("Error on CH Payout():", err.Error())
	}
	fmt.Println("8.2 SUCCESS -- ทอนเหรียญจาก hopper =", value)
	return nil
}

// coinShortage() เมื่อเหรียญไม่พอจะให้ยกเลิกการขาย โดยทำอะไรบ้าง... 1. คายธนบัตร 2. คืนเหรียญที่รับมา
func (pm *Payment) coinShortage() error {
	fmt.Println("NO -> 9 รับธนบัตรรึเปล่า")
	if pm.billEscrow != 0 { // ถ้ามียอดรับล่าสุดเป็นธนบัตร (ที่ถูกพักไว้)
		fmt.Println("YES -> 9.1 ถ้ารับด้วยธนบัตรให้คายธนบัตรคืนลูกค้า -- สั่งคายธนบัตร")
		err := BA.Take(false) // คายธนบัตร (Reject)
		if err != nil {
			return err
		}
		fmt.Println("SUCCESS -- คายธนบัตรเมื่อเหรียญใน hopper ไม่พอทอน PM.total=", PM.total)
	}
	fmt.Println("No -> 9.2 รับมาด้วยเหรียญ -- ให้คืนเหรียญตามจำนวนที่รับมา")
	err := CH.PayoutByCash(pm.coin)
	if err != nil {
		return err
		log.Println("Error on CH Payout():", err.Error())
	}
	return nil
}

func (pm *Payment) reset() {
	pm.total = 0 // เคลียร์ยอดเงินค้างให้หมด
	pm.bill = 0
	pm.billEscrow = 0
	pm.coin = 0
	pm.remain = 0
}
