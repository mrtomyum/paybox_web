package ctrl

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/static"
	"os"
)

func Router() *gin.Engine {
	r := gin.Default()
	pwd, _ := os.Getwd()
	// for Static HTML template

	//r.LoadHTMLGlob(pwd + "/view3/**/*.tpl")
	r.Static("/html", pwd+"/view3/html")
	r.Static("/js", pwd+"/view3/public/js")
	r.Static("/css", pwd+"/view3/public/css")
	r.Static("/img", pwd+"/view3/public/img")
	r.Static("/json", pwd+"/view3/public/json")
	r.Use(static.Serve("/", static.LocalFile("view3", true)))

	// WebService endpoint for web UI
	r.GET("/menu", GetMenu)
	r.GET("/menu/:id/", GetItemsByMenuId)
	r.GET("/item/:id", GetItemById)
	r.POST("/sale", NewSale)

	// WebSocket endpoint for web UI
	r.GET("/web", func(c *gin.Context) {
		ServWeb(c.Writer, c.Request)
	})

	return r
}

