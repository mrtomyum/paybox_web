package controller

import (
	"github.com/gin-gonic/gin"
	"log"
)

type msgStatus int

const (
	OK msgStatus = iota
	ERROR
	FAIL
)

type Msg struct {
	Topic  string      `json:"topic"`
	Status msgStatus   `json:"status"`
	Data   interface{} `json:"data"`
}

func GetDeviceIndexPage(ctx *gin.Context) {
	log.Print("GetDeviceIndexPage")
	ctx.HTML(200, "device", gin.H{"Title": "Hello"})
}
