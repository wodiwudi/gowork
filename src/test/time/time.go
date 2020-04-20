package main

import (
	"fmt"
	"time"
)

func main() {
	//当前时间
	now := time.Now()
	//当前时间戳(到秒)
	fmt.Println(now) //2019-10-31 21:17:52.943764 +0800 CST m=+0.000137963
	unixNow := time.Now().Unix()
	fmt.Println(unixNow) //1572527872
	//格式化时间 记住612345
	FormatNow := time.Now().Format("2006/1/2 3:04:05")
	fmt.Println(FormatNow) //2019/10/31 9:32:16
	//任意时间戳转化为格式化时间
	strTime := time.Unix(1572527872, 0).Format("2006/01/02 03 04 05")
	fmt.Println(strTime) //2019/10/31 09 17 52
	//任意格式化时间转化为时间戳
	DateTime := time.Date(2018, 5, 30, 15, 24, 59, 0, time.Local)
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", t.Local())
	//UnixDate := DateTime.Unix()
	fmt.Println()
	fmt.Println(DateTime) //1527665099
	fmt.Println(t)        //1527665099

	//输出今天是周几
	fmt.Println(now.Weekday().String())
}
