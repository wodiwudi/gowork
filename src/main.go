package main

import (
	_ "net/http" // 引入一个标准库
)

func init() {
	println("main.init.2")
}
func main() {
	test()
}
func init() {
	println("main.init.1")
}
