package main

import "github.com/gin-gonic/gin"

func main() {
	// 创建带有默认中间件的路由
	ginEngine := gin.Default()
	// 声明路由
	ginEngine.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//启动服务， 并监听在 0.0.0.0:8080
	ginEngine.Run()
}
