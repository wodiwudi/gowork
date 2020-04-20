package main

import "github.com/gin-gonic/gin"

//泛绑定 用*加任意非空字符来引导所有 /user下所有内容进到此handler
func main() {
	r := gin.Default()
	r.GET("/user/*1", func(c *gin.Context) {
		c.String(200, "user generation")
	})
	r.Run()
}
