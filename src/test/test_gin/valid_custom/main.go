package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//gtfield 大于某个字段值
type booking struct {
	Check_in  time.Time `form:"check_in" time_format:"2006-01-02" binding:"required"`
	Check_out time.Time `form:"check_out" time_format:"2006-01-02" binding:"required,gtfield=Check_in"`
}

func main() {
	r := gin.Default()
	r.GET("/book", func(c *gin.Context) {
		var book booking
		if err := c.ShouldBind(&book); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"ERROR": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "ok", "data": book})
	})

	r.Run()
}
