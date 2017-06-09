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

	//r.LoadHTMLGlob(pwd + "/view2/**/*.tpl")
	r.Static("/html", pwd+"/view2/html")
	r.Static("/js", pwd+"/view2/public/js")
	r.Static("/css", pwd+"/view2/public/css")
	r.Static("/img", pwd+"/view2/public/img")
	r.Static("/json", pwd+"/view2/public/json")
	r.Use(static.Serve("/", static.LocalFile("view2", true)))

	// WebService endpoint for web UI
	r.GET("/menu", GetMenu)
	r.GET("/menu/:id/", GetItemsByMenuId)
	r.GET("/item/:id", GetItemById)
	r.POST("/sale", NewSale)

	coin := r.Group("/coin")
	{
		coin.GET("/count", GetCoinCount)
		coin.POST("/count", SetCoinCount)
		coin.GET("/empty", EmptyCoin)
		coin.POST("/payout", PayoutCoin)
	}


	// WebSocket endpoint for web UI
	r.GET("/web", func(c *gin.Context) {
		ServWeb(c.Writer, c.Request)
	})

	return r
}

