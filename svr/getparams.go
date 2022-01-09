package svr

import (
	"encoding/json"
	"net/http"

	"ginDemo/dao"

	"github.com/gin-gonic/gin"
)

// UserInfo .
type UserInfo struct {
	Name string `json:"name"`
	Pwd  string `json:"pwd"`
}

func GetRouteParams() {
	r := gin.Default()

	r.GET("/query", func(ctx *gin.Context) {
		username, ok := ctx.GetQuery("username")
		if !ok {
			username = "默认"
		}
		ctx.JSON(http.StatusOK, gin.H{
			"username": username,
		})
	})

	r.Run()
}

func GetFormParams() {
	r := gin.Default()

	r.POST("/userInfo", func(ctx *gin.Context) {
		name := ctx.DefaultPostForm("name", "默认用户") // 获取不到值 有默认值
		age := ctx.PostForm("age")
		a, _ := dao.MustToInt32(age)
		ctx.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  a,
		})
	})

	r.Run()
}

func GetJsonParams() {
	r := gin.Default()

	r.POST("/json", func(ctx *gin.Context) {
		jsonData, err := ctx.GetRawData()
		if err != nil {
			ctx.HTML(http.StatusBadGateway, "../model/error.html", gin.H{
				"title": "error",
			})
		}
		var m map[string]interface{}
		_ = json.Unmarshal(jsonData, &m)
		ctx.JSON(http.StatusOK, m)
	})

	r.Run()
}

func ShouldBindTest() {
	r := gin.Default()
	r.POST("/user", func(ctx *gin.Context) {
		var user UserInfo
		err := ctx.ShouldBind(&user)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"massage": "error",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"status":   "ok",
			"userName": user.Name,
			"userPwd":  user.Pwd,
		})
	})
	r.Run()
}
