package svr

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func test1(ctx *gin.Context) {
	fmt.Println("test1 middle ware")
	ctx.Set("name", "xxx")
	ctx.Next() // 调用后续函数
	//ctx.Abort() // 阻止调用后续函数
}

func test2() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("test2 middle ware")
		name, _ := ctx.Get("name")
		fmt.Printf("name(%s)\n", name)
		ctx.Next()
	}
}

// Middleware 中间件
func Middleware() {
	r := gin.Default()
	r.Use(test1)
	r.GET("/middle", test2())

	r.Run()
}
