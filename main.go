package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("form.html")
	r.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(200, "form.html", nil)
	})

	//s.HelloWord()
	//s.GetRouteParams()
	//s.GetFormParams()
	//s.GetJsonParams()
	//s.ShouldBindTest()
	//s.Middleware()

	r.Run()
}
