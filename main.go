package main

import (
	"github.com/gin-gonic/gin"
	c "github.com/mrtomyum/paybox_terminal/controller"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("view/*.tpl")
	r.Static("css", "view/css")
	r.Static("img", "view/img")
	r.Static("js", "view/js")
	app := c.Router(r)
	app.Run(":8080")
}
