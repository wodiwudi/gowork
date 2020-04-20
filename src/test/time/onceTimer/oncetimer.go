package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	select {
	case <-time.After(time.Second * 5): // 5秒后执行 time   time.After返回只读channel
		// case <- time.After(周期):
		fmt.Println("after")
	}
	duration := time.Since(start)
	fmt.Println("运行时间：", duration)
}
