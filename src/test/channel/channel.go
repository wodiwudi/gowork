package main

import (
	"fmt"
	"time"
)

func worker(c chan int) {
	for n := range c { //range如果有内容会输出，如果发送方close则会break
		fmt.Printf("%d\n", n)
	}
}

func CreateWorker() chan<- int { //返回一个channel
	c := make(chan int)
	go func() { //创建一个goroutine并返回
		for { //10个goroutine一直循环等待接收
			cvalue := <-c
			fmt.Printf("value:%c\n", cvalue)
		}
	}()
	return c
}

func TestWorker() {
	var chans [10]chan<- int
	for i := 0; i < 10; i++ {
		chans[i] = CreateWorker() //循环创造10个goroutine并装进channel数组
	}
	for i := 0; i < 10; i++ {
		//添加数据
		chans[i] <- 'a' + i
	}
}
func bufferedChannel() {
	c := make(chan int, 3) //创建带缓冲区的channel
	go worker(c)
	c <- 1   //如果没有缓冲区会报deadlock
	c <- 2   //因为channel必须是两方的goroutine之间的通信，必须有人接收
	c <- 3   //而频繁的切换goroutine，即使是协程也会更加耗资源。所以加上缓冲区
	close(c) //关闭channel，发送方关闭之后
}

func main() {
	//TestWorker()
	bufferedChannel()
	time.Sleep(time.Second) //停一秒让所有创造的goroutine接收到数据并输出

}
