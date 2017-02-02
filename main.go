package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mrtomyum/paybox_terminal/ctrl"
)

func main() {

	r := gin.Default()
	app := ctrl.Router(r)
	// Dial to Device WS server
	go ctrl.CallDev()
	app.Run(":8888")
	//app.RunTLS(
	//	":8088",
	//	"api.nava.work.crt",
	//	"nava.work.key",
	//)
}
