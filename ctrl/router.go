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

	// for production use PWD path
	//r.LoadHTMLGlob(pwd + "/view/**/*.tpl")
	//r.Static("/html", pwd+"/view/html")
	//r.Static("/js", pwd+"/view/public/js")
	//r.Static("/css", pwd+"/view/public/css")
	//r.Static("/img", pwd+"/view/public/img")
	//r.Static("/json", pwd+"/view/public/json")
	//r.Use(static.Serve("/", static.LocalFile("view", true)))

	// for Test View2
	r.LoadHTMLGlob(pwd + "/view2/**/*.tpl")
	r.Static("/html", pwd+"/view2/html")
	r.Static("/js", pwd+"/view2/public/js")
	r.Static("/css", pwd+"/view2/public/css")
	r.Static("/img", pwd+"/view2/public/img")
	r.Static("/json", pwd+"/view2/public/json")
	r.Use(static.Serve("/", static.LocalFile("view2", true)))

	//// for view2
	//r.LoadHTMLGlob(pwd + "/view/**/*.tpl")
	//r.Static("/html", pwd+"/view/html")
	//r.Static("/js", pwd+"/view/public/js")
	//r.Static("/css", pwd+"/view/public/css")
	//r.Static("/img", pwd+"/view/public/img")
	//r.Static("/json", pwd+"/view/public/json")
	//r.Use(static.Serve("/", static.LocalFile("view", true)))

	// WebService endpoint for web UI
	r.GET("/menu", GetMenu)
	r.GET("/menu/:id/", GetItemsByMenuId)
	r.GET("/item/:id", GetItemById)
	r.POST("/sale", NewSale)

	// WebSocket endpoint for web UI
	r.GET("/web", func(c *gin.Context) {
		ServWeb(c.Writer, c.Request)
	})
	// เปิดบริการ WebSocket ให้โปรแกรมเว็บทดสอบ เรียกตอบเข้ามา
	//r.GET("/dev", func(c *gin.Context) {
	//	ServHw(c.Writer, c.Request)
	//})


	return r
}

