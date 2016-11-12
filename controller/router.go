package controller

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) *gin.Engine {
	r.GET("/", MenuIndex)
	r.GET("/item/:id", GetItemByMenu)
	r.POST("/order", PostOrder)
	return r
}

