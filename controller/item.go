package controller

import (
	"github.com/mrtomyum/paybox_terminal/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetItemByMenu(ctx *gin.Context) {
	var i model.Item
	id := ctx.Param("id")
	menuId, _ := strconv.ParseInt(id, 10, 64)
	items, err := i.GetIndex(menuId)
	if err != nil {
		ctx.HTML(http.StatusNotFound, "error.tpl", err.Error())
	}
	ctx.HTML(http.StatusOK, "items.tpl", items)
}
