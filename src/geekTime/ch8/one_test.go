package ch8

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

//仅需任务完成实例 ，例如从谷歌百度夸克等浏览器同时搜索词条，谁搜到返回最快的
func runTask(i int) string {
	time.Sleep(time.Microsecond * 10)
	s := fmt.Sprintf("it is come from %d", i)
	return s
}

func FirstResponse() string {
	//同时运行数
	numOfRunner := 10
	//这里要创建buffer channel 否则会造成多个协程堵塞，造成内存溢出 (oom)
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	return <-ch
}

func TestOne(t *testing.T) {
	t.Log("before : ", runtime.NumGoroutine()) //当前协程数
	t.Log(FirstResponse())
	time.Sleep(time.Second * 1)               //停止一秒 让所有协程结束。
	t.Log("after : ", runtime.NumGoroutine()) //如果不为buffer channel  算主进程有11个协程 其中9个都堵塞了
}
