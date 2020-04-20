package main

import (
	"fmt"
	"log"
	"test/queue"
)

func main() {
	myqueue := queue.Queue{1} //初始值
	myqueue.Push(2)
	myqueue.Push("3")
	myqueue.Push(4.4)
	for i := 0; i < 5; i++ {
		if res, err := myqueue.Pop(); err != nil {
			log.Fatalln(err)
		} else {
			fmt.Printf("%T,%v\n", res, res)
		}
	} // 0 1 2 3 4 queue is empty
}
