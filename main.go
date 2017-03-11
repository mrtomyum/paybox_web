package main

import (
	"github.com/mrtomyum/paybox_web/ctrl"

	"github.com/mrtomyum/paybox_web/model"
)

func main() {
	// Todo: Check NTP Server and adjust RTC
	app := ctrl.Router()
	// Dial to Device WS server
	go ctrl.CallDev()
	model.CA.Stop()
	model.BA.Stop()
	app.Run(":8888")
	//app.RunTLS(
	//	":8088",
	//	"api.nava.work.crt",
	//	"nava.work.key",
	//)
}
