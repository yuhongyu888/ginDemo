package main

import (
	s "ginDemo/svr"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("form.html")
	r.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(200, "form.html", nil)
	})

	s.HelloWord()
	s.GetRouteParams()
	s.GetFormParams()
	s.GetJsonParams()
	s.ShouldBindTest()

	r.Run()
}
