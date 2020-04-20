package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()                    //生成一个gin的默认实例
	r.GET("/ping", func(c *gin.Context) { //定义路由并在回调函数中规定返回内容
		c.JSON(200, gin.H{ //响应为json
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
