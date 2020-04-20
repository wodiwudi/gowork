package main

import "C"
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//验证绑定的结构体是否符合规范
type person struct {
	Age  int    `form:"age" binding:"required,gt=10"`
	Name string `form:"name" binding:"required"`
}

func main() {
	r := gin.Default()
	r.GET("/struct", func(c *gin.Context) {
		var person person
		err := c.ShouldBind(&person)
		if err != nil {
			c.String(http.StatusInternalServerError, "error:%v", err.Error())
			c.Abort()
			return
		}
		c.String(http.StatusOK, "%+v", person)
	})

	r.Run()
}
