package ctrl

import (
	"github.com/gin-gonic/gin"
	"github.com/mrtomyum/paybox_web/model"
	"net/http"
	"log"
)

func GetMenu(ctx *gin.Context) {
	//checkHW()
	model.BA.Stop()
	model.CA.Stop()
	var menu model.Menu
	langs, err := menu.Index()
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusNotFound, err.Error())
	}
	//// Debug =====
	//s := &model.Sale{}
	//err = model.P.Print(s)
	//if err != nil {
	//	log.Println("Error Print.")
	//}
	// Debug =====
	ctx.JSON(http.StatusOK, langs)
}

//func checkHW() {
// check HW_SERVICE websocket connected?

//model.H.Hw.Conn.SetCloseHandler(func())
// if?
//{
//	time.Sleep(1 * time.Second)
//	log.Println("HW_SERVICE is not connected!!")
//	m := model.Message{
//		Device:"web",
//		Type:"event",
//		Command:"aleart",
//		Data: "HW_SERVICE is not connected!!",
//	}
//	model.H.Web.Send <- m
//}

//}