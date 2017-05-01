package ctrl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mrtomyum/paybox_web/model"
	"log"
	"net/http"
	"time"
)

// Order ทำการบันทึกรับชำระเงิน โดยตรวจสอบการ ทอนเงิน บันทึกลง SqLite
// และส่งข้อมูล Order Post ขึ้น Cloud แต่หาก Network Down Order.completed = false
func NewSale(c *gin.Context) {
	// รับคำสั่งขายจาก Web ผ่าน JSON REST
	fmt.Println("NewSale() start")
	s := new(model.Sale)
	if c.Bind(s) != nil {
		c.JSON(http.StatusOK, gin.H{"command": "bind_sale_data", "result": "error", "data": s})
		log.Println("Error JSON from Web client.")
	}
	t := time.Now()
	s.Created = &t
	s.Type = "take_home"
	fmt.Printf("NewSale()รับค่า Sale จาก web->sale= %v saleSub = %v\n", s, &s.SaleSubs)
	//s.HostId = model.MB.MachineId()

	// Payment
	fmt.Println("Sale.Total = ", s.Total)
	err := model.PM.New(s)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusConflict, gin.H{"command": "payment", "result": "error", "message": err.Error()})
		return
	}

	// พิมพ์ตั๋ว และใบเสร็จ
	err = model.P.PrintTicket(s)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusConflict, gin.H{"command": "print", "result": "error", "message": err.Error()})
	}

	//model.P.PrintTest(data)

	// ถ้า Net IsNetOnline และ Post สำเร็จ ให้บันทึก SQL sale.completed = true
	fmt.Println("Save ยอดขายลง Local Storage")
	err = s.Save()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusConflict, gin.H{"command": "save", "result": "error", "message": err.Error()})
	}

	fmt.Println("Post ยอดขายขึ้น Cloud -> sale.Post()")
	err = s.Post()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusConflict, gin.H{"command": "post", "result": "error", "message": err.Error()})
	}

	// Reset Payment data.
	//model.PM.Reset()

	c.JSON(http.StatusOK, gin.H{"command": "sale", "result": "success", "data": s})
	fmt.Println("NewSale() COMPLETED, sale = ", s)
	s.Reset()
}
