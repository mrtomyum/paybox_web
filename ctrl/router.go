package ctrl

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/static"
)

func Router(r *gin.Engine) *gin.Engine {
	// for Static HTML template
	r.LoadHTMLGlob("view/**/*.tpl")

	// for test run disable this if production
	//r.Static("/html", "view/html")
	//r.Static("/js", "view/public/js")
	//r.Static("/css", "view/public/css")
	//r.Static("/img", "view/public/img")
	//r.Static("/json", "view/public/json")
	//r.Use(static.Serve("/", static.LocalFile("view", true)))

	// Absolute path static file for deploy production.
	r.Static("/js", "/opt/paybox/web_service/view/public/js")
	r.Static("/css", "/opt/paybox/web_service/view/public/css")
	r.Static("/img", "/opt/paybox/web_service/view/public/img")
	r.Static("/json", "/opt/paybox/web_service/view/public/json")
	r.Use(static.Serve("/", static.LocalFile("/opt/paybox/web_service/view", true)))

	// WebService endpoint for web UI
	r.GET("/menu", GetMenu)
	r.GET("/menu/:id/", GetItemsByMenuId)
	r.GET("/item/:id", GetItemById)
	r.POST("/sale", NewSale)

	// WebSocket endpoint for web UI
	r.GET("/web", func(c *gin.Context) {
		ServWeb(c.Writer, c.Request)
	})
	//r.GET("/dev", func(c *gin.Context) {
	//	ServDev(c.Writer, c.Request)
	//})


	return r
}

