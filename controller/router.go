package controller

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) *gin.Engine {
	r.GET("/", GetMenuIndex)
	r.GET("/item/:id", GetItemByMenu)
	r.POST("/order", PostNewOrderSub)
	r.DELETE("/order", DeleteOrder)
	r.DELETE("/order/item/:line", DeleteOrderItem)

	return r
}

