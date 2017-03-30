package ctrl

import (
	"github.com/gin-gonic/gin"
	"github.com/mrtomyum/paybox_web/model"
	"net/http"
	"log"
)

func GetMenu(ctx *gin.Context) {
	checkHW()
	model.BA.Stop()
	model.CA.Stop()
	var menu model.Menu
	langs, err := menu.Index()
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusNotFound, err.Error())
	}
	ctx.JSON(http.StatusOK, langs)
}

func checkHW() {
	//model.H.Hw.Conn.SetCloseHandler(func())

	{
		//time.Sleep(1 * time.Second)
		//log.Println("HW_SERVICE is not connected!!")
		//m := model.Message{
		//Device:"web",
		//Type:"event",
		//Command:"alert",
		//Data: "HW_SERVICE is not connected!!",
		//}
		//model.H.Web.Send <- m
	}

}