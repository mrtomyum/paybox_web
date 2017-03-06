package ctrl

import (
	"github.com/mrtomyum/paybox_web/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"log"
	"fmt"
)

func GetItemById(ctx *gin.Context) {
	var item model.Item
	strId := ctx.Param("id")
	id, _ := strconv.ParseInt(strId, 10, 64)
	err := item.Get(id)
	if err != nil {
		log.Println("GetItemById error:", err)
		ctx.HTML(http.StatusNotFound, "error.tpl", err)
		//ctx.HTML(http.StatusNotFound, "error.tpl", err.Error())
	}
	fmt.Println("Return JSON:", item)
	ctx.JSON(http.StatusOK, item)
}

func GetItemsByMenuId(ctx *gin.Context) {
	fmt.Println("call GetItemsByMenuId")
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Println("error:", err)
	}
	var item model.Item
	langs, err := item.ByMenuId(id)
	if err != nil {
		ctx.HTML(http.StatusNotFound, "error.tpl", err.Error())
	}
	ctx.JSON(http.StatusOK, langs)
}