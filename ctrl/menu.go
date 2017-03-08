package ctrl

import (
	"github.com/gin-gonic/gin"
	"github.com/mrtomyum/paybox_web/model"
	"net/http"
)

func GetMenu(ctx *gin.Context) {
	var menu model.Menu
	langs, err := menu.Index()
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
	}
	ctx.JSON(http.StatusOK, langs)
}
