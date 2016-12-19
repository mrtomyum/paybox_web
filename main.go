package main

import (
	"github.com/gin-gonic/gin"
	c "github.com/mrtomyum/paybox_terminal/controller"
	"github.com/mrtomyum/paybox_terminal/model"
)

func main() {
	hub := model.MyHub
	go hub.Start()
	r := gin.Default()
	app := c.Router(r)
	app.Run(":8888")

}
