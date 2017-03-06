package main

import (
	"github.com/mrtomyum/paybox_terminal/ctrl"

)

func main() {

	app := ctrl.Router()
	// Dial to Device WS server
	go ctrl.CallDev()
	app.Run(":8888")
	//app.RunTLS(
	//	":8088",
	//	"api.nava.work.crt",
	//	"nava.work.key",
	//)
}
