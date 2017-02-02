package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mrtomyum/paybox_terminal/ctrl"
	"log"
)

func main() {
	// Dial to Device WS server
	err := ctrl.CallDev()
	if err != nil {
		//log.Println("Error call Device Websocket:", err)
		log.Fatal("Error dial:", err)
	}

	r := gin.Default()
	app := ctrl.Router(r)
	app.Run(":8888")
	//app.RunTLS(
	//	":8088",
	//	"api.nava.work.crt",
	//	"nava.work.key",
	//)
}
