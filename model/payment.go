package model

import (
	"errors"
	"fmt"
	"log"
)

var ErrCoinShortage error = errors.New("Not enough coin to change// ไม่มีเหรียญพอทอน")

// Payment คือยอดเงินพัก ยังไม่ได้รับชำระ
type Payment struct {
	coin       float64 // มูลค่าเหรียญที่รับมาทั้งหมด แต่ยังไม่ได้รับชำระ
	bill       float64 // มูลค่าธนบัตรที่รับมาทั้งหมด แต่ยังไม่ได้รับชำระ
	billEscrow float64 // มูลค่าธนบัตรที่พักอยู่ในตัวรับธนบัตร BA
	total      float64 // มูลค่าเงินพักทั้งหมด
	remain     float64 // เงินคงค้างชำระ
	change     float64 // เงินทอน
	//billCh     chan *Message
	billCh     chan float64
	//coinCh     chan *Message
	coinCh     chan float64
	cancelCh   chan bool
	isOpen     bool // เปิดรับชำระหรือยัง?
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

//func (pm *Payment) Order(s *Socket) {
//	BA.Start()
//	CA.Start()
//	msg := &Message{}
//	for {
//		select {
//		case msg = <-pm.billCh:
//			//pm.New()
//		case msg = <-pm.coinCh:
//		}
//	}
//}

// init() ทำการรีเซ็ทค่าที่ควรถูกตั้งใหม่ทุกครั้งที่สร้าง Payment ใหม่ขึ้นมา
func (pm *Payment) init(s *Sale) {
	pm.Reset()
	pm.remain = s.Total
	pm.isOpen = true
	AB = &AcceptedBill{
		B20:   true,
		B50:   true,
		B100:  true,
		B500:  true,
		B1000: true,
	}
	// เปิดการรับชำระธนบัตร และ เหรียญ (Set Inhibit)
	fmt.Println("func New() -- 1. Start Payment device:")
	BA.Start()
	CA.Start()
}

// *Payment New() ทำหน้าที่จัดการกระบวนการรับเงิน ทอนเงิน ให้สมบูรณ์
func (pm *Payment) New(s *Sale) error {
	fmt.Printf("s.Total = %v, pm.total = %v\n", s.Total, pm.total)
	CheckSocketConnected()
	if s.Total == 0 {
		return errors.New("Sale Total is 0 cannot do payment.")
	}
	pm.init(s)
	//defer close(pm.billCh)
	//defer close(pm.coinCh)
	//defer close(pm.cancelCh)

	sp := new(SalePay)
	s.SalePay = sp // ล้างข้อมูลเดิมถ้ามี

	msg := &Message{}

	// หากธนบัตร หรือเหรียญที่ชำระยังมีมูลค่าน้อยกว่ายอดขาย (Payment < Sale)
	// ระบบจะ Take เงิน, สะสมยอดรับชำระ และแจ้ง "onhand" ให้ UI วนไปจนกว่าจะได้ยอด Payment >= Sale
	for pm.total < s.Total {
		// Todo: ให้ Cancel() ทำการส่ง pm.billCh เข้ามาด้วย msg.command: "cancel" แล้วทำการตรวจ if cancel ให้ break ออกจาก loop
		pm.adjAcceptedBill(s)    // ปรับยอดด้างชำระ เพื่อกำหนดชนิดธนบัตรที่ยอมรับได้
		pm.displayAcceptedBill() // displayAcceptedBill() ส่งรายการธนบัตรที่รับได้ไปแสดงบนหน้าจอ
		pm.sendOnHand(H.Web)     // ส่งยอดรับเงินปัจจุบันให้ UI

		fmt.Println("1. Waiting payment form BA or CA")
		var value float64
		select {
		case cancel := <-pm.cancelCh:
			fmt.Println("case <-pm.cancelCh return...")
			if cancel {
				return errors.New("cancel")
			}
			//case <-pm.cancelCh:
			//	fmt.Println("case <-pm.cancelCh return...")
			//	return errors.New("cancel")
		case value = <-pm.billCh:
			//value = msg.Data.(float64)
			fmt.Printf("3. Bill Accepted: pm.total= %v sale.total= %v pm.remain= %v msg = %v\n", pm.total, s.Total, pm.remain, msg)
			pm.billEscrow = value
			fmt.Println("pm.billEscrow:", pm.billEscrow)
			if pm.billEscrow == 0 { // ถ้าไม่มีธนบัตรพักอยู่
				return ErrNoBillEscrow
			}
			if pm.isAcceptedBill(value) { // ถ้ายอมรับธนบัตรราคานี้
				BA.Take() // ให้เก็บธนบัตรลงถัง
				fmt.Printf("4. เก็บธนบัตรสำเร็จ: pm.total= %v sale.total= %v pm.remain= %v\n", pm.total, s.Total, pm.remain)
			} else { // ถ้าไม่รับ
				BA.Reject() // ให้คายทิ้ง และล้างยอดรับเงิน/ ยอดค้างรับกลับไปเริ่มต้น รอรับเงินใหม่
				pm.billEscrow = 0
			}

		case value = <-pm.coinCh:
			fmt.Printf("3. Coin Accepted: pm.total= %v sale.total= %v pm.remain= %v\n", pm.total, s.Total, pm.remain)
		}
		// บันทึกประเภทเหรียญและธนบัตรที่รับมาลง s.SalePay
		fmt.Println("value=", value)
		err := sp.Add(value)
		if err != nil {
			return err
		}
		fmt.Printf("5. footer: pm.total= %v sale.total= %v pm.remain= %v\n", pm.total, s.Total, pm.remain)
	}

	// ปิดการรับชำระที่ อุปกรณ์
	CA.Stop()
	BA.Stop()
	fmt.Printf("6. pm.total= %v sale.total= %v pm.remain= %v\n", pm.total, s.Total, pm.remain)

	// ทอนเงิน
	pm.change = pm.total - s.Total
	fmt.Printf("pm.change = %v pm.total = %v s.Total = %v", pm.change, pm.total, s.Total)
	if pm.change > 0 { //ถ้าต้องทอนเงิน
		err := pm.doChange(pm.change)
		if err != nil {
			return err
		}
	}

	// Update ยอดรับเงิน และทอนเงินให้ Sale{}
	s.Pay = pm.total
	s.Change = pm.change

	// ส่งยอด Onhand ให้ UI
	m := &Message{
		Device:  "web",
		Command: "payment",
		Type:    "event",
		Data:    "success",
	}
	H.Web.Send <- m
	return nil
}

// Cancel คืนเงินจากทุก Device โดยตรวจสอบเงิน Escrow ใน bill Acceptor ด้วยถ้ามีให้คืนเงิน
func (pm *Payment) Cancel() {
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
	//fmt.Printf("pm.billEscrow: %v pm.total: %v ", pm.billEscrow, pm.total)
	change := pm.total - pm.billEscrow
	switch {
	case pm.billEscrow != 0: //มีธนบัตร
		// สั่งให้ BillAcceptor คืนเงินที่พักไว้ ซึ่งจะคืนได้เพียงใบล่าสุด
		BA.Reject() // คายธนบัตร
		fallthrough
	case change != 0: // CoinHopper สั่งให้จ่ายเหรียญที่คงค้างตามยอดคงเหลือ PM.coin ออกด้านหน้า
		err := CH.PayoutByCash(change)
		if err != nil {
			log.Println(err.Error())
		}
	}
	BA.Stop()
	CA.Stop()
	pm.cancelCh <- true
	pm.Reset()
}

// OnHand ส่งค่าเงินพัก Escrow ไว้กลับไปให้ web
func (pm *Payment) sendOnHand(s *Socket) {
	fmt.Println("method *Host.sendOnHand()... pm.total =", pm.total)
	s.Msg.Device = "web"
	s.Msg.Command = "onhand"
	s.Msg.Result = true
	s.Msg.Type = "event"
	s.Msg.Data = pm.total
	s.Send <- s.Msg
}

func (pm *Payment) adjAcceptedBill(s *Sale) {
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
func (pm *Payment) isAcceptedBill(v float64) bool {
	fmt.Println("4. ตรวจว่าเป็นธนบัตรที่รับได้หรือไม่?")
	switch v {
	case 20.0:
		if !AB.B20 {
			return false
		}
	case 50.0:
		if !AB.B50 {
			return false
		}
	case 100.0:
		if !AB.B100 {
			return false
		}
	case 500:
		if !AB.B500 {
			return false
		}
	case 1000:
		if !AB.B1000 {
			return false
		}
	default:
		fmt.Printf("มูลค่า BillEscrow = %v ไม่เข้าเงื่อนไข\n", pm.billEscrow)
	}
	fmt.Println("PM.billEscrow =", pm.billEscrow, "AcceptedBill = ", AB)
	return true
}

// ปรับยอดเงินจากธนบัตรที่รับมา
func (pm *Payment) addBill(v float64) {
	// อัพเดตยอดเงินสดในตู้ด้วย
	CB.bill += pm.billEscrow  // เพิ่มยอดธนบัตรในถังธนบัตร
	CB.total += pm.billEscrow // เพิ่มยอดรวมของ CashBox
	pm.bill += pm.billEscrow
	pm.total += pm.billEscrow
	pm.remain -= pm.billEscrow
	pm.billEscrow = 0 // ล้างยอดเงินพัก
}

func (pm *Payment) rejectBill(v float64) {
	BA.Reject()
	pm.billEscrow = 0 // ล้างยอดเงินพัก
}

func (pm *Payment) addCoin(value float64) {
	pm.coin += value
	pm.total += value
	pm.remain -= value
	CB.hopper += value
	CB.total += value
	fmt.Println("coin billCh =", pm.coin, "pm total=", pm.total)
}

// refund() เมื่อต้องการพิมพ์ใบคืนเงิน ในกรณีเงินทอนไม่พอ
func (pm *Payment) refund(total, billEscrow float64) error {
	BA.Reject() //คายธนบัตร
	// Todo: Print ใบคืนเงิน (Refund) ตามยอดเงินคงเหลือ
	rf := total - billEscrow
	err := P.makeRefund(rf) //ยังไม่เสร็จ
	if err != nil {
		return err
	}
	return nil
}

func (pm *Payment) doChange(value float64) error {
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
	go CH.PayoutByCash(value) // Todo: เพิ่มกลไกวิเคราะห์เงินทอน แล้วสั่งทอนเป็นเหรียญ เพื่อป้องกันเหรียญหมด
	//if err != nil {
	//	return err
	//	log.Println("Error on CH Payout():", err.Error())
	//}
	//fmt.Println("8.2 SUCCESS -- ทอนเหรียญจาก hopper =", value)
	return nil
}

// coinShortage() เมื่อเหรียญไม่พอจะให้ยกเลิกการขาย โดยทำอะไรบ้าง... 1. คายธนบัตร 2. คืนเหรียญที่รับมา
func (pm *Payment) coinShortage() error {
	fmt.Println("NO -> 9 รับธนบัตรรึเปล่า")
	if pm.billEscrow != 0 { // ถ้ามียอดรับล่าสุดเป็นธนบัตร (ที่ถูกพักไว้)
		fmt.Println("YES -> 9.1 ถ้ารับด้วยธนบัตรให้คายธนบัตรคืนลูกค้า -- สั่งคายธนบัตร")
		BA.Reject() // คายธนบัตร (Reject)
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

func (pm *Payment) Reset() {
	log.Println("Reset Payment: ล้างยอดรับชำระ")
	pm.total = 0 // เคลียร์ยอดเงินค้างให้หมด
	pm.bill = 0
	pm.billEscrow = 0
	pm.coin = 0
	pm.remain = 0
	pm.change = 0
	pm.isOpen = false
	resetChannel(PM.billCh)
	resetChannel(BA.Send)
	resetChannel(CA.Send)
	resetChannel(CH.Send)
	resetChannel(MB.Send)
	resetChannel(P.Send)
	// ส่งยอดที่ล้างแล้วให้ WebUI
	pm.sendOnHand(H.Web)
}

func (pm *Payment) Change() float64 {
	return pm.change
}

func CheckSocketConnected() {
	// ตรวจสอบ WebSocket Connection?
	switch {
	case H.Hw == nil:
		log.Println("HW_SERVICE websocket ยังไม่ได้เชื่อมต่อ")
	case H.Web == nil:
		log.Println("WEB UI websocket ยังไม่ได้เชื่อมต่อ")
	}
}
