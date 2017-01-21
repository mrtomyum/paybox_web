package ctrl

import (
	"fmt"
	"github.com/mrtomyum/paybox_terminal/model"
	"log"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Order ทำการบันทึกรับชำระเงิน โดยตรวจสอบการ ทอนเงิน บันทึกลง SqLite
// และส่งข้อมูล Order Post ขึ้น Cloud แต่หาก Network Down Order.completed = false
func NewSale(c *gin.Context) {
	// รับคำสั่งจาก Web ผ่าน JSON REST
	fmt.Println("NewSale() start")
	sale := &model.Sale{}
	if c.Bind(sale) != nil {
		c.JSON(http.StatusBadRequest, sale)
		log.Println("Error JSON from Web client.")
	}
	fmt.Printf("[NewSale()] รับค่า Order จาก web-> sale= %v\n", sale)

	// DisplayAcceptedBill() จากยอดขาย ส่งรายการธนบัตรที่รับได้ไปแสดงบนหน้าจอ
	DisplayAcceptedBill()

	// เริ่มการรับชำระที่อุปกรณ์ทุกตัว (Set Inhibit)
	// BillAcceptor:BA
	model.BA.Start()
	// CoinAccptor:CA
	model.CA.Start()

	// Total Pay >= Total Sale? หากธนบัตร หรือเหรียญที่ชำระยังมีมูลค่าน้อยกว่ายอดขาย (Payment < Sale)
	// กรณีธนบัตรระบบจะ Take เงิน และจะ สะสมยอดรับชำระ และส่ง command: "onhand" เป็น event กลับตลอดเวลาจนกว่าจะได้ยอด Payment >= Sale

	// หากรายการสุดท้ายชำระเป็นธนบัตร ระบบจะยังไม่ Take เงิน โดยตรวจสอบว่ามีเงินทอนเพียงพอหรือไม่?
	// หากมากพอ ระบบจะทอนเงิน หากไม่พอ ระบบจะ Reject ธนบัตรใบล่าสุดนี้คืน และส่ง Message แจ้งเตือนให้เปลี่ยนธนบัตร หรือเหรียญ (ข้อความจะเปลี่ยนตามภาษาที่เลือก)
	// หากรายการสุดท้ายชำระเป็นธนบัตร ระบบจะยังไม่ Take เงิน โดยตรวจสอบว่ามีเงินทอนเพียงพอหรือไม่? หากมากพอ ระบบจะทอนเงิน
	// หากไม่พอ ระบบจะ Reject ธนบัตรใบล่าสุดนี้คืน และส่ง Message แจ้งเตือนให้เปลี่ยนธนบัตร หรือเหรียญ (ข้อความจะเปลี่ยนตามภาษาที่เลือก)
	if model.OH.Coin > sale.Total {
		err := model.BA.Take(false)
		if err != nil {
			c.JSON(http.StatusConflict, gin.H{"result":"error", "message":err.Error()})
		}
		model.OH.Total = - model.OH.Bill
		model.OH.Bill = 0
	}

	// กินธนบัตรที่พักไว้ *ระวัง! ถ้า Dev client ยังไม่เปิดคอนเนคชั่นจะ runtime error: invalid memory address or nil pointer derefere
	err := model.BA.Take(true)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"result":"error", "message":err.Error()})
		log.Println("Error on Bill_Acceptor Take():", err.Error())
	}

	// ทอนเงินจาก CoinHopper ถ้ามี
	if model.OH.Total > sale.Total {
		change := model.OH.Total - sale.Total
		err = model.CH.PayoutByCash(change) // Todo: เพิ่มกลไกวิเคราะห์เงินทอน แล้วสั่งทอนเป็นเหรียญ เพื่อป้องกันเหรียญหมด
		if err != nil {
			c.JSON(http.StatusConflict, gin.H{"result":"error", "message":err.Error()})
			log.Println("Error on CH Payout():", err.Error())
		}
	}

	// อัพเดตยอดเงินสดในตู้ด้วย
	model.CB.Bill = + model.OH.Bill
	model.CB.Total = + model.OH.Bill
	model.OH.Total = - model.OH.Bill
	model.OH.Bill = 0

	// พิมพ์ตั๋ว และใบเสร็จ
	model.P.Print(sale)

	// ส่งยอดเงินพักในมือให้ web client ล้างยอดเงิน
	model.H.OnHand(model.H.Web)

	// เช็คสถานะ Network และ Server ว่า IsNetOnline อยู่หรือไม่?
	if !model.H.IsNetOnline {
		fmt.Println("Offline => Save sale to disk")
	}
	fmt.Println("sale.Post()")
	sale.Post()

	// ถ้า Net IsNetOnline และ Post สำเร็จ ให้บันทึก SQL sale.completed = true
	fmt.Println("sale.Save()")
	err = sale.Save()
	if err != nil {
		log.Println("Error in sale.Save() =>", err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"result":       "success",
		"total_escrow": model.OH.Total,
	})
	fmt.Println("NewSale() COMPLETED, sale = ", sale)
}

func DisplayAcceptedBill() {
	// Check MinAcceptedBill500 & 1000
	m := &model.Message{
		Command:"accepted_bill",
		Type:   "event",
		Data:   model.AB,
	}
	model.H.Web.Send <- m
}