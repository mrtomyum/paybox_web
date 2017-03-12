package ctrl

import (
	"fmt"
	"github.com/mrtomyum/paybox_web/model"
	"log"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Order ทำการบันทึกรับชำระเงิน โดยตรวจสอบการ ทอนเงิน บันทึกลง SqLite
// และส่งข้อมูล Order Post ขึ้น Cloud แต่หาก Network Down Order.completed = false
func NewSale(c *gin.Context) {

	// รับคำสั่งจาก Web ผ่าน JSON REST
	fmt.Println("NewSale() start")
	sale := model.Sale{}
	if c.Bind(sale) != nil {
		c.JSON(http.StatusOK, gin.H{"command": "bind_sale_data", "result": "error", "data": sale, })
		log.Println("Error JSON from Web client.")
	}
	fmt.Printf("[NewSale()] รับค่า Sale จาก web->sale= %v\n", sale)

	err := sale.New()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusConflict, gin.H{"command": "post", "result": "error", "message": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"command": "sale", "result": "success", "data": sale, })
	fmt.Println("NewSale() COMPLETED, sale = ", sale)
}
