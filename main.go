package main

import (
	"github.com/mrtomyum/paybox_web/ctrl"
	"fmt"
)

func main() {
	app := ctrl.Router()
	fmt.Println("1")

	// Dial to HW_SERVICE

	go ctrl.OpenSocket()
	fmt.Println("2")

	// Run Web Server
	app.Run(":8888")

	//app.RunTLS(
	//	":8088",
	//	"api.nava.work.crt",
	//	"nava.work.key",
	//)
}
