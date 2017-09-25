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
	r.Static("/html", pwd+"/view5/html")
	r.Static("/js", pwd+"/view5/static/js")
	r.Static("/css", pwd+"/view5/static/css")
	r.Static("/img", pwd+"/view5/static/img")
	r.Static("/fonts", pwd+"/view5/static/fonts")
	r.Static("/voice", pwd+"/view5/static/voice")
	r.Use(static.Serve("/", static.LocalFile("view5", true)))

	// WebService endpoint for web UI
	r.GET("/menu", GetMenu)
	r.GET("/menu/:id/", GetItemsByMenuId)
	r.GET("/item/:id", GetItemById)
	r.POST("/sale", NewSale)
	r.POST("/saletest", NewSale)
	r.OPTIONS("/saletest", NewSale)


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

