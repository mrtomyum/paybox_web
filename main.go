package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mrtomyum/paybox_terminal/model"
	"github.com/mrtomyum/paybox_terminal/ctrl"
)

func main() {
	go model.H.Run()
	r := gin.Default()
	app := ctrl.Router(r)
	app.Run(":8088")
	//app.RunTLS(
	//	":8088",
	//	"api.nava.work.crt",
	//	"nava.work.key",
	//)
}


