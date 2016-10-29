package controller

import (
	"github.com/mrtomyum/paybox_terminal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetItemIndex(ctx *gin.Context) {
	var i model.Item
	items, err := i.GetIndex(DB)
	if err != nil {
		ctx.HTML(http.StatusNotFound, "error.tpl", err.Error())
	}
	ctx.HTML(http.StatusOK, "item.tpl", items)
}
