package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mrtomyum/paybox_terminal/model"
	"net/http"
	"strconv"
)

func GetIndex(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "index", nil)
}

func GetMenu(ctx *gin.Context) {
	var m model.Menu
	menus, err := m.Index()
	if err != nil {
		ctx.HTML(http.StatusNotFound, "error.tpl", err.Error())
	}
	ctx.HTML(http.StatusOK, "list", menus)

}

func GetItemByMenuId(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	var i model.Item
	items, err := i.FindById(id)
	if err != nil {
		ctx.HTML(http.StatusNotFound, "error.tpl", err.Error())
	}
	ctx.HTML(http.StatusOK, "items", items)
}