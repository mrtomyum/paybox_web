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
	fmt.Printf("[NewSale()] รับค่า Sale จาก web->sale= %v\n", sale)

	// DisplayAcceptedBill() จากยอดขาย ส่งรายการธนบัตรที่รับได้ไปแสดงบนหน้าจอ
	DisplayAcceptedBill()

	// Payment
	err := model.PM.Pay(sale)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"command": "payment", "result":"error", "message":err.Error()})
	}

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
		"onhand":       model.PM.Total,
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
	fmt.Println("Send message to Web = ", m)
	model.H.Web.Send <- m
}