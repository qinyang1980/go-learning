package web

import "github.com/gin-gonic/gin"

// 启动gin的web服务
func StartHttpServer() {
	r := gin.Default()

	r.GET("/test1", func(c *gin.Context) {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "bad request",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
