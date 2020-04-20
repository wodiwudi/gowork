package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/test", func(c *gin.Context) {
		first := c.PostForm("first")                            //获取post数据
		second := c.DefaultPostForm("second", "default_second") //有默认值的获取post数据
		c.String(http.StatusOK, "%s,%s", first, second)
		/*bytes, err := ioutil.ReadAll(c.Request.Body) //ioutil.ReadAll获取body内容为byte
		if err != nil {
			c.String(http.StatusBadRequest, err.Error()) //错误
			c.Abort()                                    //中断
		}
		c.String(http.StatusOK, string(bytes)) //转成string输出*/
	})
	r.Run()
}
