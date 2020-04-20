package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//获取get参数
func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		first := c.Query("first")               //不带默认值的获取get参数
		second := c.DefaultQuery("second", "2") //如果没有该参数，则默认为2
		c.String(http.StatusOK, "%s,%s", first, second)
	})
	r.Run()
}
