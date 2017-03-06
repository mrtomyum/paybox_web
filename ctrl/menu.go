package ctrl

import (
	"github.com/gin-gonic/gin"
	"github.com/mrtomyum/paybox_web/model"
	"net/http"
)

func GetIndex(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "index", nil)
}

func GetMenu(ctx *gin.Context) {
	var menu model.Menu
	langs, err := menu.Index()
	if err != nil {
		ctx.HTML(http.StatusNotFound, "error.tpl", err.Error())
	}
	ctx.JSON(http.StatusOK, langs)
}
