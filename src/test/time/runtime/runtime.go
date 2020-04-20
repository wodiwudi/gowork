package main

import (
	"fmt"
	"time"
)

/*
time.Hour 表示1小时
time.Minute 表示1分钟
time.Second 表示1秒
time.Millisecond	表示1毫秒
time.Microsecond	表示1微妙
time.Nanosecond	表示1纳秒
*/

func StartCac() {
	t1 := time.Now()

	for i := 0; i < 10000; i++ {
		time.Sleep(2 * time.Nanosecond) //休眠2纳秒
		// continue
	}
	elapsed := time.Since(t1) //time.Since() 从设定时间段到这行代码运行的时间 单位：毫秒
	fmt.Println("App elapsed: ", elapsed)
}

func main() {
	StartCac() //App elapsed:  5.846651ms

}
