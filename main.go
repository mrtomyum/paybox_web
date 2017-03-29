package main

import (
	"github.com/mrtomyum/paybox_web/ctrl"
	"time"
)

func main() {
	// Todo: Check NTP Server and adjust RTC
	app := ctrl.Router()

	// Dial to Device WS server
	go ctrl.CallHw()
	time.Sleep(3 * time.Second)

	app.Run(":8888")
	//app.RunTLS(
	//	":8088",
	//	"api.nava.work.crt",
	//	"nava.work.key",
	//)
}
