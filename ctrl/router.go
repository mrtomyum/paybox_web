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

	// for test run disable this if production
	r.LoadHTMLGlob(pwd + "/view/**/*.tpl")
	r.Static("/html", pwd+"/view/html")
	r.Static("/js", pwd+"/view/public/js")
	r.Static("/css", pwd+"/view/public/css")
	r.Static("/img", pwd+"/view/public/img")
	r.Static("/json", pwd+"/view/public/json")
	r.Use(static.Serve("/", static.LocalFile("view", true)))

	// Absolute path static file for deploy production.
	//r.LoadHTMLGlob("/opt/paybox/web_service/view/**/*.tpl")
	//r.Static("/js", "/opt/paybox/web_service/view/public/js")
	//r.Static("/css", "/opt/paybox/web_service/view/public/css")
	//r.Static("/img", "/opt/paybox/web_service/view/public/img")
	//r.Static("/json", "/opt/paybox/web_service/view/public/json")
	//r.Use(static.Serve("/", static.LocalFile("/opt/paybox/web_service/view", true)))

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
	r.GET("/dev", func(c *gin.Context) {
		ServDev(c.Writer, c.Request)
	})


	return r
}

