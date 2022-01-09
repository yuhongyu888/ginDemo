package svr

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloWord() {
	r := gin.Default()

	r.GET("/index", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "hello word",
		})
	})

	r.POST("/post", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "post",
		})
	})

	r.Run()
}
