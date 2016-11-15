package main

import (
	"github.com/gin-gonic/gin"
	c "github.com/mrtomyum/paybox_terminal/controller"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("view/**/*.tpl")
	r.Static("/public", "./view/public")
	app := c.Router(r)
	app.Run(":8080")
}


