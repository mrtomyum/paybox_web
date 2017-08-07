package ctrl

import (
	"github.com/gin-gonic/gin"
	"github.com/mrtomyum/paybox_web/model"
	"net/http"
	"log"
)

func GetMenu(ctx *gin.Context) {
	//model.BA.Stop()
	//model.CA.Stop()
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,Token")
	var menu model.Menu
	langs, err := menu.Index()
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusNotFound, err.Error())
	}
	ctx.JSON(http.StatusOK, langs)
}
