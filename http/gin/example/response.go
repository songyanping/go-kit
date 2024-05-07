package main

import (
	"github.com/gin-gonic/gin"
	"github.com/songyanping/go-kit/http/gin/response"
)

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	r.GET("/welcome", func(c *gin.Context) {
		// 从请求的查询参数中获取name值，默认值为"Guest"
		name := c.DefaultQuery("name", "Guest")
		response.OkWithMessage(name, c)
	})
	r.Run()
}
