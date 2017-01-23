package model

import (
	"fmt"
	"log"
)

// Payment คือยอดเงินพัก ยังไม่ได้รับชำระ
type Payment struct {
	Coin  float64 // มูลค่าเหรียญพัก ที่ยังไม่ได้รับชำระ
	Bill  float64 // มูลค่าธนบัตรที่พักอยู่ในเครื่องรับธนบัตร
	Card  float64 // มูลค่าบัตรเครดิตที่รับชำระแล้ว
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

// AcceptedBill ระบุค่ายอดขายขั้นต่ำที่ยอมรับธนบัตรแต่ละขนาด 0 = ไม่จำกัด
type AcceptedBill struct {
	THB20   int `json:"thb_20"`
	THB50   int `json:"thb_50"`
	THB100  int `json:"thb_100"`
	THB500  int `json:"thb_500"`
	THB1000 int `json:"thb_1000"`
}

func (oh *Payment) Pay(sale *Sale) error {
	// เปิดการรับชำระธนบัตร และ เหรียญ (Set Inhibit)
	fmt.Printf("func Pay() -- \n1. Start inhibit device BA, CA \n")
	BA.Start()
	CA.Start()

	// ให้รอจนกว่าจะได้รับเงิน จาก BA หรือ CA
	fmt.Println("2. Waiting payment form BA or CA")
	m := <-PM.Send
	fmt.Println("Received Money:", m.Data)

	// หากธนบัตร หรือเหรียญที่ชำระยังมีมูลค่าน้อยกว่ายอดขาย (Payment < Sale)
	// ระบบจะ Take เงิน และจะสะสมยอดรับชำระ และส่ง command: "onhand" เป็น event กลับตลอดเวลาจนกว่าจะได้ยอด Payment >= Sale
	for PM.Total < sale.Total {
		if m.Device == "bill_acc" { // เฉพาะธนบัตรต้องสั่ง Take ก่อน
			// กินธนบัตรที่พักไว้ *ระวัง! ถ้า Dev client ยังไม่เปิดคอนเนคชั่นจะ runtime error: invalid memory address or nil pointer derefere
			err := BA.Take(true)
			if err != nil {
				return err
			}
		}
	}

	// เมื่อชำระเงินครบหรือเกิน ตรวจว่ามีเหรียญพอทอนหรือไม่?
	change := PM.Total - sale.Total
	// หากรายการสุดท้ายชำระเป็นธนบัตร ระบบจะยังไม่ Take เงิน โดยตรวจสอบว่ามีเงินทอนเพียงพอหรือไม่? หากมากพอ ระบบจะทอนเงิน
	// หากไม่พอ ระบบจะ Reject ธนบัตรใบล่าสุดนี้คืน และส่ง Message แจ้งเตือนให้เปลี่ยนธนบัตร หรือเหรียญ (ข้อความจะเปลี่ยนตามภาษาที่เลือก)
	if CB.Hopper > change {
		err := BA.Take(false)
		if err != nil {
			return err
		}
		PM.Total = - PM.Bill
		PM.Bill = 0
	}

	// ทอนเงินจาก CoinHopper ถ้ามี
	if PM.Total > sale.Total {

		err := CH.PayoutByCash(change) // Todo: เพิ่มกลไกวิเคราะห์เงินทอน แล้วสั่งทอนเป็นเหรียญ เพื่อป้องกันเหรียญหมด
		if err != nil {
			return err
			log.Println("Error on CH Payout():", err.Error())
		}
	}

	// ปิดการรับชำระที่อุปกรณ์
	BA.Stop()
	CA.Stop()
	return nil
}
