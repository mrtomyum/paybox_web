package main

import (
	"github.com/mrtomyum/paybox_web/ctrl"
)

func main() {
	// Todo: Check NTP Server and adjust RTC and Test Hardware status
	app := ctrl.Router()

	// Dial to HW_SERVICE

	go ctrl.OpenSocket()

	// Run Web Server
	app.Run(":8888")

	//app.RunTLS(
	//	":8088",
	//	"api.nava.work.crt",
	//	"nava.work.key",
	//)
}
