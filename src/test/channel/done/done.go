package main

import (
	"fmt"
	"sync"
)

type worker struct { //创建结构体放两个channel
	in   chan int
	done func()
}

func doworker(w worker) {
	for n := range w.in { //10个goroutine一直循环等待接收
		fmt.Printf("value:%c\n", n)
		w.done()
	}
}

func CreateWorker(wg *sync.WaitGroup) worker { //返回一个channel
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doworker(w)
	return w
}

func TestWorker() {
	var wg sync.WaitGroup
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = CreateWorker(&wg) //循环创造10个goroutine并装进channel数组
	}
	wg.Add(10)
	for i, worker := range workers {
		//添加数据
		worker.in <- 'a' + i
	}
	wg.Wait()
}

func main() {
	TestWorker()
}
