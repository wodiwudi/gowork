package main

import "github.com/gin-gonic/gin"

//静态文件
func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.Run() //在编译器运行会找不到静态文件   用go build -o router_static && ./router_static
}
