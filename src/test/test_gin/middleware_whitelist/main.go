package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//自定义白名单中间件
func IPAuthMiddleware() gin.HandlerFunc { //返回值必须为gin中的HandlerFunc
	return func(c *gin.Context) {
		iplist := []string{ //ip白名单列表
			"127.0.0.2",
		}
		flag := false
		clientIP := c.ClientIP() //获取客户端ip
		for _, host := range iplist {
			if host == clientIP {
				flag = true
				break
			}
		}
		if !flag {
			c.String(http.StatusUnauthorized, "%s,not in iplist", clientIP)
			c.Abort()
		}
	}
}

func main() {
	r := gin.Default()
	r.Use(IPAuthMiddleware()) //自己定义一个中间件
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "123")
	})
	r.Run()
}
