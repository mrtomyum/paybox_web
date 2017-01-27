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
	sale := model.S
	if c.Bind(sale) != nil {
		c.JSON(http.StatusOK, gin.H{"command":"bind_sale_data", "result": "error", "data": sale, })
		log.Println("Error JSON from Web client.")
	}
	fmt.Printf("[NewSale()] รับค่า Sale จาก web->sale= %v\n", sale)

	// Payment
	err := model.PM.Pay(sale)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"command": "payment", "result":"error", "message":err.Error()})
	}

	// พิมพ์ตั๋ว และใบเสร็จ
	err = model.P.Print(sale)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"command": "print", "result":"error", "message":err.Error()})
	}

	// ส่งยอดเงินพักในมือให้ web client ล้างยอดเงิน
	model.PM.OnHand(model.H.Web)

	fmt.Println("Post ยอดขายขึ้น Cloud -> sale.Post()")
	err = sale.Post()
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"command": "post", "result":"error", "message":err.Error()})
	}

	// ถ้า Net IsNetOnline และ Post สำเร็จ ให้บันทึก SQL sale.completed = true
	fmt.Println("Save ยอดขายลง Local Storage")
	err = sale.Save()
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"command": "save", "result":"error", "message":err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"command":"sale", "result": "success", "data": sale, })
	fmt.Println("NewSale() COMPLETED, sale = ", sale)
}
