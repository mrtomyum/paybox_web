package main

import (
	"github.com/gin-gonic/gin"
	c "github.com/mrtomyum/paybox_terminal/controller"
	_ "net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("view/**/*.tpl")
	r.Static("/html", "./view/html")
	r.Static("/public", "./view/public")
	r.Static("/js", "./view/public/js")
	r.Static("/css", "./view/public/css")
	r.Static("/img", "./view/public/img")
	r.Static("/json", "./view/public/json")
	app := c.Router(r)
	app.Run(":8888")
}


