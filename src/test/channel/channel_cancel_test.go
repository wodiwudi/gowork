package main

import (
	"fmt"
	"testing"
	"time"
)

//任务的取消
//关闭通道函数，有任意值进入则true
func iscancelChan(ch chan interface{}) bool {
	select {
	case <-ch:
		return true
	default:
		return false
	}
}
func cancelChan(ch chan interface{}) {
	ch <- 1 //随意发送任意值，只关闭第一个接收到数据的协程
}
func cancelChan2(ch chan interface{}) {
	close(ch) //关闭通道，是广播形式的，所有基于此的协程均关闭
}

func TestCancelChan(t *testing.T) {
	ch := make(chan interface{})
	for i := 0; i < 5; i++ {
		//创造5个协程
		go func(num int, ch chan interface{}) {
			for {
				if iscancelChan(ch) {
					break
				}
				time.Sleep(time.Second * 1)
			}
			fmt.Printf("%d is done !", num)
			time.Sleep(time.Second * 1)
		}(i, ch)
	}
	cancelChan2(ch)
	time.Sleep(time.Second * 1)
}
