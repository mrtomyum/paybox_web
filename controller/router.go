package controller

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) *gin.Engine {
	r.GET("/", MenuIndex)
	r.GET("/items", GetItemIndex)
	return r
}

