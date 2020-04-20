package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func main() {
	//创建一个文件用来记录日志和错误
	f, _ := os.Create("./gin.log")
	gin.DefaultWriter = io.MultiWriter(f)      //记录日志
	gin.DefaultErrorWriter = io.MultiWriter(f) //记录错误
	//gin.Default中相当于new一个引擎然后加载各种中间件
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery()) //使用logger中间件和recovery中间件
	r.GET("/test", func(c *gin.Context) {
		//panic(1)  如果没后recovery中间件 panic会终止服务
		name := c.DefaultQuery("name", "default_name")
		c.String(http.StatusOK, "%s", name)
	})

	r.Run()
}
