package main

import "C"
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//传参绑定结构体
//结果结构体
type person struct {
	Name    string    `form:"name"`
	Address string    `form:"address"`
	Birth   time.Time `form:"birth" time_format:"2006-01-02"` //日期格式
}

func bindStruct(c *gin.Context) {
	var person person
	error := c.ShouldBind(&person) //将内容绑定到结构体，传入结构体指针
	if error != nil {
		c.String(http.StatusOK, "error %v", error.Error())
	}
	c.String(http.StatusOK, "%v", person)
}

func main() {
	r := gin.Default()
	r.GET("struct", bindStruct) //post与get都绑定到一个方法
	r.POST("struct", bindStruct)
	r.Run()
}
