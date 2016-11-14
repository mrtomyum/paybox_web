package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mrtomyum/paybox_terminal/model"
	"net/http"
)

func GetMenuIndex(ctx *gin.Context) {
	var m model.Menu
	menus, err := m.Index()
	if err != nil {
		ctx.HTML(http.StatusNotFound, "error.tpl", err.Error())
	}
	ctx.HTML(http.StatusOK, "menus.tpl", menus)
}