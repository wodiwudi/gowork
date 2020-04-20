package main

import (
	"errors"
	"fmt"
)

func tryRecover() {
	//panic仍旧会调用defer
	defer func() {
		res := recover() //recover必须在defer函数中调用才能生效
		//Type asseting断言判断类型
		if err, ok := res.(error); ok {
			fmt.Println("recover a error :", err)
		}
	}()
	panic(errors.New("a panic"))
}

func main() {
	tryRecover()
}
