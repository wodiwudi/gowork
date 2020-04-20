package main

import (
	"fmt"
	"time"
)

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(j int) { //由于协程是非抢占式的a[j]会一直加加,导致死锁 因为没有人切换协程到main函数让它sleep。
			// fmt包中的函数会不断切换线程所以不会死循环
			for {
				fmt.Println("xx")
			}
		}(i) //不可以直接用for中的i，因为在程序执行时，for循环早已经跑完10次，i= 10 ，但是goroutine慢一些
		// 会导致使用 i= 10的那个i 会报index out of range错误。而且数组下边都串位置了
	}
	time.Sleep(time.Minute)
	fmt.Println(a)
}
