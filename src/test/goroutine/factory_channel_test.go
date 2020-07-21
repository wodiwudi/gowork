package main

import (
	"fmt"
	"sync"
	"testing"
)

//将goroutine和通道绑定
type recever struct {
	sync.WaitGroup
	data chan int
}

func newRecever() *recever {
	r := &recever{
		data: make(chan int),
	}
	r.Add(1)
	go func() {
		defer r.Done()
		for x := range r.data {
			fmt.Println("recv: ", x)
		}
	}()
	return r
}

func TestFactoryChannel(t *testing.T) {
	recv := newRecever()
	recv.data <- 1
	recv.data <- 2
	close(recv.data)
	recv.Wait()
}
