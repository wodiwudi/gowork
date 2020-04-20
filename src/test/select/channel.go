package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(c chan int) {
	for n := range c {
		fmt.Printf("%d\n", n)
	}
}

func CreateWorker() chan<- int { //返回一个channel
	c := make(chan int)
	go worker(c)
	return c
}

func generator() chan int {
	out := make(chan int)

	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}

	}()
	return out
}

func testSelect() {
	var c1, c2 = generator(), generator()
	w := CreateWorker()
	for {
		select {
		case n := <-c1:
			w <- n
		case n := <-c2:
			w <- n

		}
	}

}

func main() {
	testSelect()
}
