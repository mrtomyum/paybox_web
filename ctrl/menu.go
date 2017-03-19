package ctrl

import (
	"github.com/gin-gonic/gin"
	"github.com/mrtomyum/paybox_web/model"
	"net/http"
	"log"
)

func GetMenu(ctx *gin.Context) {
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
