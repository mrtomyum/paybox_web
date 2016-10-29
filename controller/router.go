package controller

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) *gin.Engine {
	r.GET("/", MenuIndex)
	item := r.Group("/item")
	{
		item.GET("/", GetItemIndex)
	}
	return r
}

