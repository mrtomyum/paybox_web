package main

import (
	"github.com/gin-gonic/gin"
	c "github.com/mrtomyum/paybox_terminal/controller"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("template/**/*.tpl")
	r.Static("/public", "./public")
	//r.Static("css", "view/css")
	//r.Static("img", "view/img")
	//r.Static("js", "view/js")
	app := c.Router(r)
	app.Run(":8080")
}
